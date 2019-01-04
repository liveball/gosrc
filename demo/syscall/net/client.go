package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

const (
	family  = syscall.AF_INET
	sotype  = syscall.SOCK_STREAM | syscall.SOCK_STREAM
	proto   = 0
	backlog = 128

	EPOLLET        = 1 << 31
	MaxEpollEvents = 32
)

func main() {
	var (
		err error
		fd  int
	)
	fd, err = syscall.Socket(family, sotype, proto)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		err = syscall.Close(fd)
		if err != nil {
			fmt.Printf("error(%v)\n", err)
		}
	}()

	addr := syscall.SockaddrInet4{Port: 9000}
	copy(addr.Addr[:], net.ParseIP("127.0.0.1").To4())

	err = syscall.Connect(fd, &addr)
	if err != nil {
		fmt.Printf("syscall.Connect error(%v)\n", err)
		os.Exit(1)
	}
	fmt.Println("connect")
	send(fd)
	syscall.Close(fd)
}

func send(fd int) {
	for {
		syscall.Write(fd, []byte(time.Now().Format("15:04:05")))

		time.Sleep(1 * time.Second)

		buf := make([]byte, 1024)
		n, err := syscall.Read(fd, buf)
		if n <= 0 {
			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK {
				fmt.Println("we have read all")
				break
			} else {
				fmt.Println("syscall.Read: ", err)
				break
			}
		}

	}
}
