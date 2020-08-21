package main

import (
	"context"
	"testing"
)

func TestNewPool_Acquire_Release(t *testing.T) {
	p := NewPool(5)

	i := 5

	for i > 0 {
		i--

		c, err := p.Acquire(context.Background())
		if err != nil {
			t.Error(err)
			return
		}

		p.Release(c)

		n, err := c.Write([]byte("hello"))
		if err != nil {
			t.Error(err)
			return
		}

		t.Log(n)
	}
}
