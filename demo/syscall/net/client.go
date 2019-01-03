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
	sotype  = syscall.SOCK_STREAM
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

	// if err = syscall.SetNonblock(fd, true); err != nil {
	// 	fmt.Printf("syscall.SetNonblock error(%v)\n", err)
	// 	os.Exit(1)
	// }

	addr := syscall.SockaddrInet4{Port: 9000}
	copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())

	err = syscall.Connect(fd, &addr)
	if err != nil {
		fmt.Printf("syscall.Connect error(%v)\n", err)
		os.Exit(1)
	}
	fmt.Println("connect")

	syscall.Write(fd, []byte(time.Now().Format("15:04:05\n")))
	// poller(fd)
}

func poller(fd int) {
	var (
		err  error
		epfd int
	)
	epfd, err = syscall.EpollCreate1(0)
	if err != nil {
		fmt.Println("syscall.EpollCreate1: ", err)
		os.Exit(1)
	}
	defer syscall.Close(epfd)

	var event syscall.EpollEvent
	var events [MaxEpollEvents]syscall.EpollEvent
	event.Events = syscall.EPOLLIN | EPOLLET

	if err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &event); err != nil {
		fmt.Println("syscall.EpollCtl: ", err)
		os.Exit(1)
	}
	for {
		n, err := syscall.EpollWait(epfd, events[:], -1)
		if err != nil {
			fmt.Println("syscall.EpollWait: ", err)
			break
		}

		for i := 0; i < n; i++ {
			ev := events[i]
			if ev.Events&(syscall.EPOLLERR|syscall.EPOLLHUP|syscall.EPOLLIN) == 0 {
				fmt.Printf("epoll error at event: %d ", ev.Events)
				continue
			}

			go readWrite(int(ev.Fd))
		}
	}
}

func readWrite(fd int) {
	defer syscall.Close(fd)
	for {
		buf := make([]byte, 0, 1024)
		n, err := syscall.Read(fd, buf)
		if err != nil {
			fmt.Printf("read syscall.Read error(%v)\n", err)
			break
		}
		println("client readWrite...", n)
		if n == 0 {
			break
		}
		syscall.Write(fd, []byte(time.Now().Format("15:04:05\n")))
		time.Sleep(1 * time.Second)
	}
}
