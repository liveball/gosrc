package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalln(err)
		return
	}

	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Println(err)
			return
		}
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	fmt.Println("close")
	conn.Close()
	<-done
	log.Println("exit")
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println(err)
	}
}
