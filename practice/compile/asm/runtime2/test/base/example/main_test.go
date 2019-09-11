package main

import (
	"testing"

	"gosrc/practice/compile/asm/runtime2/test/base"
)

var Result int64

func BenchmarkAddNative(b *testing.B) {
	var r int64
	for i := 0; i < b.N; i++ {
		r = int64(i) + int64(i)
	}
	Result = r
}

func BenchmarkAddAsm(b *testing.B) {
	var r int64
	for i := 0; i < b.N; i++ {
		r = base.Add(int64(i), int64(i))
	}
	Result = r
}