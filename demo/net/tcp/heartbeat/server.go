package heartbeat

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

//Msg for conn
type Msg struct {
	data []byte
}

//Conn for heartbeat conn
type Conn struct {
	conn         net.Conn
	readTimeout  time.Duration
	writeTimeout time.Duration

	readChan  chan *Msg
	writeChan chan *Msg
	closeChan chan struct{}
	isClosed  bool

	mux sync.Mutex
}

// Server listen.
func Server() {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8888, Zone: ""})
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	c := &Conn{
		conn:         conn,
		readTimeout:  time.Duration(time.Millisecond * 3000),
		writeTimeout: time.Duration(time.Millisecond * 3000),

		readChan:  make(chan *Msg, 5),
		writeChan: make(chan *Msg, 5),
		closeChan: make(chan struct{}),
	}

	go c.procChanLoop()
	go c.readConnLoop()
	go c.writeConnLoop()
}

func (c *Conn) procChanLoop() {
	go func() {
		for {
			if err := c.write([]byte("heartbeat\n")); err != nil {
				log.Fatalf("c.write heartbeat error(%v)", err)
				c.close()
				return
			}

			time.Sleep(time.Minute * 1)
		}
	}()

	for {
		m, err := c.read()
		if err != nil {
			log.Fatalf("c.read error(%v)", err)
			c.close()
			break
		}

		err = c.write(m.data)
		if err != nil {
			log.Fatalf("c.write error(%v)", err)
			c.close()
			break
		}

	}

}

func (c *Conn) write(bs []byte) error {
	select {
	case c.writeChan <- &Msg{data: bs}:
	case <-c.closeChan:
		return errors.New("conn closed")
	}
	return nil
}

func (c *Conn) read() (*Msg, error) {
	select {
	case m := <-c.readChan:
		return m, nil
	case <-c.closeChan:
		return nil, errors.New("conn closed")
	}
}

func (c *Conn) readConnLoop() {

	var buf *bufio.Reader
	for {
		c.conn.SetReadDeadline(time.Now().Add(c.readTimeout)) //设置读超时
		buf = bufio.NewReader(c.conn)
		bs, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatalf("c.conn.Read error(%v)", err)
			c.close()
			return
		}

		fmt.Print("read from client:", string(bs))

		req := &Msg{
			data: bs,
		}

		select {
		case c.readChan <- req:
		case <-c.closeChan:
			return
		}
	}
}

func (c *Conn) writeConnLoop() {
	for {
		select {
		case m := <-c.writeChan:
			c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
			if _, err := c.conn.Write(m.data); err != nil {
				log.Fatalf("c.conn.Write error(%v)", err)
				c.close()
				return
			}
		case <-c.closeChan:
			return
		}

	}
}

func (c *Conn) close() {
	c.conn.Close()

	c.mux.Lock()
	if !c.isClosed {
		c.isClosed = true
		close(c.closeChan)
	}
	c.mux.Unlock()
}
