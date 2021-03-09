package main

import "testing"

func BenchmarkSafeXORBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		safeXORBytes(g, e, f)
	}
}

func BenchmarkFastXORWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastXORWords(g, e, f)
	}
}

func TestBit(t *testing.T) {
	a := uint(1)
	b := uint(1) >> 63

	c := 4 << (^uint(1) >> 63)

	t.Log(a, b, c)
}
