package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	// for i := 0; i < 10; i++ {
	// 	go dial()
	// }

	dial()
	// select {}
}

func dial() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("net.Dial error(%v)\n", err)
		return
	}
	defer conn.Close()

	for {
		//read
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("conn error(%v)\n", err)
			// conn.Close()
			return
		}
		fmt.Print("Message from server: " + message)

		// write
		conn.Write([]byte("hello\n"))

		time.Sleep(1 * time.Second)
	}
}

// mustCopy(os.Stdout, conn)
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
