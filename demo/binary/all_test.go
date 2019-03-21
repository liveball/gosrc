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
