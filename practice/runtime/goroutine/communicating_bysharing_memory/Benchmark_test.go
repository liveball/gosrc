package main

import "testing"

func Benchmark_addByShareMemory(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		addByShareMemory(10)
	}
}

//goos: darwin
//goarch: amd64
//pkg: gosrc/practice/runtime/goroutine/communicating_bysharing_memory
//Benchmark_addByShareMemory 	 1000000	      1937 ns/op
//PASS

func Benchmark_addByShareCommunicate(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		addByShareCommunicate(10)
	}
}

//goos: darwin
//goarch: amd64
//pkg: gosrc/practice/runtime/goroutine/communicating_bysharing_memory
//Benchmark_addByShareCommunicate 	  500000	      2836 ns/op
//PASS
