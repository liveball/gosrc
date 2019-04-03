package main

import (
	"math/rand"
	"testing"
)

func BenchmarkMakeZeroCapSlice(b *testing.B) {
	a := make([]int, 0)
	for i := 0; i < b.N; i++ {
		r := rand.Intn(100)
		for j := 0; j < r; j++ {
			a = append(a, j)
		}
		a = make([]int, 0)
	}
}

func BenchmarkMakeCapSlice(b *testing.B) {
	a := make([]int, 0)
	for i := 0; i < b.N; i++ {
		r := rand.Intn(100)
		for j := 0; j < r; j++ {
			a = append(a, j)
		}
		a = make([]int, 0, 100)
		a = a[:0] //at count=1000000 77.554856ms
	}
}

func BenchmarkNewSlice(b *testing.B) {
	a := make([]int, 0)
	for i := 0; i < b.N; i++ {
		r := rand.Intn(100)
		for j := 0; j < r; j++ {
			a = append(a, j)
		}
		a = a[:0]
	}
}
