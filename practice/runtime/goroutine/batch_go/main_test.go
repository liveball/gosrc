package main

import "testing"

func BenchmarkBatchG(b *testing.B) {
	b.Run("batchG", func(b *testing.B) {
		batchG()
	})
}
