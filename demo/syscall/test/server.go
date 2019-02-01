package main

import (
	"go1.11.1/demo/syscall/net/epoll"
	"go1.11.1/demo/syscall/net/socket"
)

var poller *epoll.Poller

func init() {
	poller = epoll.New()
}

func main() {
	sk := socket.New("127.0.0.1", 9000, 32)
	sk.Poller = poller
	sk.Listen()
	for {
		sk.Accept()
	}
}
