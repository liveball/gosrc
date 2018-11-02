package heartbeat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//Client for dial.
func Client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	var buf *bufio.Reader
	go func() {
		for {
			conn.SetReadDeadline(time.Now().Add(time.Second * 3)) //每次读写设置超时

			buf = bufio.NewReader(conn)

			msg, err := buf.ReadBytes('\n')
			if err != nil {
				log.Fatalf("buf.ReadBytes error(%v)", err)
				return
			}

			fmt.Println("the number of bytes:", buf.Buffered())

			fmt.Print("read from server:", string(msg))

			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		for {
			conn.SetWriteDeadline(time.Now().Add(time.Second * 3)) //每次写设置超时

			conn.Write([]byte(time.Now().Format("15:04:05\n")))

			time.Sleep(time.Second * 1)
		}
	}()

	for {
		conn.Write([]byte("heartbeat\n"))
		time.Sleep(time.Minute * 1)
	}
}
