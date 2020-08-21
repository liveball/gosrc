package main

import (
	"net"
	"sync"
)

type Pool struct {
	mu   sync.Mutex
	cond sync.Cond

	numConns, limit int
	idle            []net.Conn
}

func NewPool(limit int) *Pool {
	p := &Pool{
		limit: limit,
	}
	p.cond.L = &p.mu
	return p
}

func (p *Pool) Acquire() (net.Conn, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.idle) == 0 &&
		p.numConns >= p.limit {
		p.cond.Wait()
	}

	if len(p.idle) > 0 {
		c := p.idle[len(p.idle)-1]
		p.idle = p.idle[:len(p.idle)-1]
		return c, nil
	}

	c, err := net.Dial("tcp", "127.0.0.1:8888")
	if err == nil {
		p.numConns++
	}

	return c, err
}

func (p *Pool) Release(c net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.idle = append(p.idle, c)
	p.cond.Signal()
}

func (p *Pool) Hijack(c net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.numConns--
	p.cond.Signal()
}
