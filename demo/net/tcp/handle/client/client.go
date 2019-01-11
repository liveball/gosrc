package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {

	// for i := 0; i < 10; i++ {
	// 	go dial()
	// }

	dial()
}

func dial() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("net.Dial error(%v)\n", err)
		return
	}
	defer conn.Close()

	for {
		conn.Write([]byte(time.Now().Format("hello"))) // 如果不加 \n,服务端采用bufio.NewReader(conn).ReadString('\n') 将读不到，会阻塞
		// conn.SetReadDeadline(time.Now().Add(time.Second * 3)) //每次读设置超时

		// message, err := bufio.NewReader(conn).ReadString('\n')
		buf := make([]byte, 32)
		n, err := bufio.NewReader(conn).Read(buf)
		// n, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("client buf.ReadString error(%v)\n", err)
			return
		}
		fmt.Println("from server: " + string(buf[:n]))
		time.Sleep(1 * time.Second)
	}
}
