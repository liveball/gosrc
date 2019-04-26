package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	if _, err := conn.Write([]byte("hello\n")); err != nil {
		log.Fatalln(err)
		return
	}
}
