package main

import (
	"runtime"
	"testing"
)

func fin(v *int) {
}

func BenchmarkFinalizer(b *testing.B) {
	const Batch = 1000
	b.RunParallel(func(pb *testing.PB) {
		var data [Batch]*int
		for i := 0; i < Batch; i++ {
			data[i] = new(int)
		}
		for pb.Next() {
			for i := 0; i < Batch; i++ {
				runtime.SetFinalizer(data[i], fin)
			}
			for i := 0; i < Batch; i++ {
				runtime.SetFinalizer(data[i], nil)
			}
		}
	})
}
