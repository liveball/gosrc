package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
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
	client := flag.String("client", "client1", "which client?")
	flag.Parse() //解析输入的参数

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

	var wg sync.WaitGroup
	wg.Add(2)
	go recv(fd, &wg)
	go send(fd, &wg, *client)
	wg.Wait()
	syscall.Close(fd)
}

func recv(fd int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		buf := make([]byte, 5)
		n, err := syscall.Read(fd, buf)
		if err != nil {
			if err == syscall.EAGAIN {
				fmt.Println("we have read all")
				return
			} else {
				fmt.Println("syscall.Read: ", err)
				return
			}
		}
		if n <= 0 {
			fmt.Println("syscall.Read n==0 ")
			return
		}

		fmt.Printf("syscall.Read recv n(%d) buf(%v)\n", n, buf[:n])
		time.Sleep(1 * time.Second)
	}
}

func send(fd int, wg *sync.WaitGroup, client string) {
	defer wg.Done()
	for {
		_, err := syscall.Write(fd, []byte(client+"=>"+time.Now().Format("15:04:05")))
		if err != nil {
			fmt.Printf("write syscall.Write: fd(%d) error(%v)\n", fd, err)
			return
		}
		// fmt.Printf("write syscall.Write: n(%d)\n", n)
		time.Sleep(1 * time.Second)
	}
}
