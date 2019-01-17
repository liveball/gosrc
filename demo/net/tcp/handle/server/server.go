package main

import (
	"bufio"
	"fmt"
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
	for {
		// message, err := bufio.NewReader(conn).ReadString('\n')

		buf := make([]byte, 32)
		n, err := bufio.NewReader(conn).Read(buf)
		// n, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("server buf.ReadString error(%v)\n", err)
			return
		}
		fmt.Println("from client: " + string(buf[:n]))

		conn.Write([]byte(time.Now().Format("15:04:05"))) //如果不加 \n,客户端采用bufio.NewReader(conn).ReadString('\n') 读不到，会阻塞
		time.Sleep(1 * time.Second)
	}
}
