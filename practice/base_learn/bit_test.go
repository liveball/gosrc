package main

import "testing"

func TestBit(t *testing.T) {
	a := uint(1)
	b := uint(1) >> 63
	c := ^uint(1)
	d := ^uint(1) >> 63
	e := 4 << (^uint(1) >> 63)

	t.Log(a, b, c, d, e)
}
