package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
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

	if err = syscall.SetNonblock(fd, true); err != nil {
		fmt.Printf("syscall.SetNonblock error(%v)\n", err)
		os.Exit(1)
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	if err != nil {
		fmt.Printf("syscall.SetsockoptInt error(%v)\n", err)
		os.Exit(1)
	}

	addr := syscall.SockaddrInet4{Port: 9000}
	copy(addr.Addr[:], net.ParseIP("0.0.0.0").To4())
	err = syscall.Bind(fd, &addr)
	if err != nil {
		fmt.Printf("syscall.Bind error(%v)\n", err)
		os.Exit(1)
	}
	err = syscall.Listen(fd, backlog)
	if err != nil {
		fmt.Printf("syscall.Listen error(%v)\n", err)
		os.Exit(1)
	}

	lsa, _ := syscall.Getsockname(fd)
	fmt.Println("listen:", lsa)

	poller(fd)
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

			if int(ev.Fd) == fd { //accept
				go accept(epfd, ev)
				continue
			} else {
				go readWrite(int(ev.Fd))

			}
		}
	}
}

func accept(epfd int, event syscall.EpollEvent) {
	fd := int(event.Fd)
	defer syscall.Close(fd)
	for {
		connFd, _, err := syscall.Accept(fd)
		if connFd == -1 {
			if err == syscall.EAGAIN || err == syscall.EWOULDBLOCK {
				fmt.Println("we have processed all incoming conns")
				break
			} else {
				fmt.Println("syscall.Accept: ", err)
				break
			}
		}

		err = syscall.SetNonblock(fd, true)
		if err != nil {
			fmt.Printf("syscall.SetNonblock error(%v)\n", err)
			break
		}

		sn, _ := syscall.Getsockname(fd)
		fmt.Printf("accept new conn(%v)\n", sn)

		event.Fd = int32(connFd)
		event.Events = syscall.EPOLLIN | EPOLLET
		if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, &event); err != nil {
			fmt.Print("syscall.EpollCtl: ", connFd, err)
			os.Exit(1)
		}
	}
}

func readWrite(fd int) {
	defer syscall.Close(fd)
	for {
		buf := make([]byte, 0, 1024)
		n, err := syscall.Read(fd, buf)

		println("server readWrite...")
		if err != nil {
			fmt.Printf("read syscall.Read error(%v)\n", err)
			break
		}
		if n == 0 {
			break
		}

		fmt.Printf("read buf(%s)", buf)
		// syscall.Write(fd, buf)
		// time.Sleep(1 * time.Second)
	}
}
