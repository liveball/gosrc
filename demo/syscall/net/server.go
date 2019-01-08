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
	sotype  = syscall.O_NONBLOCK | syscall.SOCK_STREAM
	proto   = 0
	backlog = 128

	MaxEpollEvents = 32
	EPOLLET        = 0x80000000
)

func main() {
	var (
		err error
		fd  int
	)
	fd, err = syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	// fd, err = syscall.Socket(family, sotype, proto)
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
	copy(addr.Addr[:], net.ParseIP("127.0.0.1").To4())
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
	syscall.Close(fd) //close listen fd
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

	event.Fd = int32(fd)
	event.Events = syscall.EPOLLIN | syscall.EPOLLOUT | syscall.EPOLLRDHUP | EPOLLET
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
			if ev.Events == 0 {
				continue
			}

			// if (ev.Events&syscall.EPOLLERR) != 0 ||
			// 	(ev.Events&syscall.EPOLLHUP) != 0 {
			// 	fmt.Printf("epoll error ev.Events(%v)\n", ev.Events)
			// 	syscall.Close(int(ev.Fd))
			// 	continue
			// }

			// fmt.Println("read:", ev.Events, syscall.EPOLLIN, ev.Events&(syscall.EPOLLIN))
			// fmt.Println("write:", ev.Events, syscall.EPOLLOUT, ev.Events&(syscall.EPOLLOUT))

			if int(ev.Fd) == fd { //accept
				go accept(epfd, ev)
			} else if ev.Events&(syscall.EPOLLIN) != 0 { //read
				go read(epfd, ev)
			} else if ev.Events&(syscall.EPOLLOUT) != 0 { //write
				go write(epfd, ev)
			} else {
				fmt.Printf("poller continue ev(%v)\n", ev)
			}
		}
	}
}

func accept(epfd int, event syscall.EpollEvent) {
	fd := int(event.Fd)
	for {
		connFd, _, err := syscall.Accept(fd)
		if connFd == -1 {
			if err == syscall.EAGAIN {
				fmt.Println("we have processed all incoming conns")
				return
			} else {
				fmt.Println("syscall.Accept: ", err)
				return
			}
		}

		err = syscall.SetNonblock(fd, true)
		if err != nil {
			fmt.Printf("syscall.SetNonblock error(%v)\n", err)
			return
		}

		event.Fd = int32(connFd)
		event.Events = syscall.EPOLLIN | EPOLLET
		if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, connFd, &event); err != nil {
			fmt.Print("syscall.EpollCtl: ", connFd, err)
			return
		}

		fmt.Printf("accept new conn(%v)\n", connFd)
	}
}

func read(epfd int, event syscall.EpollEvent) {
	fd := int(event.Fd)
	defer syscall.Close(fd)

	for {
		buf := make([]byte, 32)
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

		fmt.Printf("syscall.Read buf(%v)\n", string(buf[:n]))

		event.Fd = int32(fd)
		event.Events = syscall.EPOLLOUT | EPOLLET
		if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_MOD, fd, &event); err != nil {
			fmt.Print("read syscall.EpollCtl: ", fd, err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func write(epfd int, event syscall.EpollEvent) {
	fd := int(event.Fd)
	defer syscall.Close(fd)

	for {
		data := []byte{1, 2, 3, 4, 5}
		// addr := syscall.SockaddrInet4{Port: 9000}
		// copy(addr.Addr[:], net.ParseIP("127.0.0.1").To4())
		// if err := syscall.Sendto(fd, data[:], 0, &addr); err != nil {
		// 	fmt.Print("write syscall.Sendto: ", fd, err)
		// 	os.Exit(1)
		// }

		_, err := syscall.Write(fd, data)
		if err != nil {
			fmt.Printf("write syscall.Write: fd(%d) error(%v)\n", fd, err)
			return
		}
		// fmt.Printf("write syscall.Write: n(%d)\n", n)

		event.Fd = int32(fd)
		event.Events = syscall.EPOLLIN | EPOLLET
		if err := syscall.EpollCtl(epfd, syscall.EPOLL_CTL_MOD, fd, &event); err != nil {
			fmt.Print("write syscall.EpollCtl: ", fd, err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
