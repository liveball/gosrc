package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("l.Accept error(%v)\n", err)
			return
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {

	defer conn.Close()
	var buf *bufio.Reader
	buf = bufio.NewReader(conn)
	for {
		//read
		bs, err := buf.ReadBytes('\n') //log agent

		if err != nil {
			if err == io.EOF {
				fmt.Printf("conn error(%v)", err)
			}
			return
		}

		if len(bs) > 0 {
			fmt.Print("read from client:", string(bs))
		}

		// write
		// conn.Write([]byte(time.Now().Format("15:04:05\n")))

		time.Sleep(1 * time.Second)
	}
}
