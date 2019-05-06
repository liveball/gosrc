package socket

import (
	"log"
	"net"
	"syscall"

	"go1.11.1/demo/syscall/net/epoll"
)

type Lisenter struct {
	Poller  *epoll.Poller
	FD      int
	IP      string
	Port    int
	Backlog int
}

func New(ip string, port, backlog int) *Lisenter {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_STREAM, 0)
	if err != nil {
		log.Fatalf("syscall.Socket error(%v)", err)
	}
	addr := syscall.SockaddrInet4{Port: port}
	copy(addr.Addr[:], net.ParseIP(ip).To4())
	err = syscall.Bind(fd, &addr)
	if err != nil {
		log.Fatalf("syscall.Bind error(%v)", err)
	}

	return &Lisenter{
		FD:      fd,
		IP:      ip,
		Port:    port,
		Backlog: backlog,
	}
}

func (c *Lisenter) Listen() {
	err := syscall.Listen(c.FD, c.Backlog)
	if err != nil {
		log.Fatalf("syscall.Listen error(%v)", err)
	}
	c.Poller.Wait(c.FD)
}

func (c *Lisenter) Accept() {
	connFd, _, err := syscall.Accept(c.Poller.Event.FD)
	if connFd == -1 {
		if err == syscall.EAGAIN {
			log.Fatalf("syscall.EAGAIN error(%v)", err)
		} else {
			log.Fatalf("syscall.Accept error(%v)", err)
		}
	}

	err = syscall.SetNonblock(c.FD, true)
	if err != nil {
		log.Fatalf("syscall.SetNonblock error(%v)", err)
		return
	}

	c.Poller.Add(connFd)
	log.Printf("accept new conn(%d)\n", connFd)
}

func (c *Lisenter) Read(buf []byte) {
	for {
		_, err := syscall.Read(c.FD, buf)

		if err != nil {
			if err == syscall.EAGAIN {
				log.Println("we have read all")
				return
			} else {
				log.Println("syscall.Read: ", err)
				return
			}
		}
		c.Poller.Modify(c.FD)
	}
}

func (c *Lisenter) Write(buf []byte) {
	for {
		_, err := syscall.Write(c.FD, buf)
		if err != nil {
			log.Printf("syscall.Write: fd(%d) error(%v)", c.FD, err)
			return
		}
		c.Poller.Modify(c.FD)
	}
}

func (c *Lisenter) Close() {
	syscall.Close(c.FD)
}
