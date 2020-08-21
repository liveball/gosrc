package main

import (
	"context"
	"fmt"
	"net"
)

type token struct{}

type Pool struct {
	sem  chan token
	idle chan net.Conn
}

func NewPool(limit int) *Pool {
	sem := make(chan token, limit)
	idle := make(chan net.Conn, limit)
	return &Pool{
		sem,
		idle,
	}
}

func (p *Pool) Acquire(ctx context.Context) (net.Conn, error) {
	select {
	case conn := <-p.idle:
		fmt.Println("idle conn")
		return conn, nil
	case p.sem <- token{}:
		conn, err := net.Dial("tcp", "127.0.0.1:8888")
		if err == nil {
			<-p.sem
		}

		fmt.Println("new conn")
		return conn, err

	case <-ctx.Done():
		return nil, ctx.Err()
	}

}

func (p *Pool) Release(c net.Conn) {
	p.idle <- c
}

func (p *Pool) Hijack(c net.Conn) {
	<-p.sem
}
