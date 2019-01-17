package main

import (
	"context"
	"fmt"
	"os"
)

type Conn interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type TCPConn struct {
	f *os.File
}

func (tc *TCPConn) Read(b []byte) (n int, err error) {
	tc.f.Seek(0, 0) //移回指针到文件开头
	n, err = tc.f.Read(b)
	// fmt.Println("read", tc.f.Fd())
	return
}

func (tc *TCPConn) Write(b []byte) (n int, err error) {
	n, err = tc.f.Write(b)
	return
}

// type Server interface {
// 	Read([]byte) (int, error)
// 	Write([]byte) (int, error)
// }

type Poller struct {
	c Conn
}

func (p *Poller) Read(c context.Context, b []byte) (n int, err error) {
	n, err = p.c.Read(b)
	return
}

func (p *Poller) Write(c context.Context, b []byte) (n int, err error) {
	n, err = p.c.Write(b)
	return
}

func main() {
	f, err := os.OpenFile("a.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
	defer f.Close()
	if err != nil {
		fmt.Printf("os.Open error (%v)\n", err)
		return
	}

	p := &Poller{c: &TCPConn{f}}
	var (
		wn, rn int
	)
	wn, err = p.Write(context.Background(), []byte("hello\n"))
	if err != nil {
		fmt.Printf("p.Write error (%v)\n", err)
		return
	}
	fmt.Println(wn)

	buf := make([]byte, 1024)
	rn, err = p.Read(context.Background(), buf)
	if err != nil {
		fmt.Printf("p.Read error (%v)\n", err)
		return
	}
	fmt.Println(string(buf[:rn]))
}
