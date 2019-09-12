package bench_atomic_mutex

import (
	"testing"
	"sync"
	"sync/atomic"
)

var (
	a,b int64
	mux sync.Mutex
)


func Benchmark_Atom(b *testing.B){
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Atom()
	}
}

//goos: darwin
//goarch: amd64
//pkg: gosrc/practice/runtime/atom/bench_atomic_mutex
//Benchmark_Atom-12    	300000000	         5.85 ns/op
//PASS

func Benchmark_Mutex(b *testing.B){
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Mutex()
	}
}

//goos: darwin
//goarch: amd64
//pkg: gosrc/practice/runtime/atom/bench_atomic_mutex
//Benchmark_Mutex-12    	100000000	        13.6 ns/op
//PASS

func Atom(){
	atomic.AddInt64(&a,1)
}

func Mutex(){
	mux.Lock()
	b++
	mux.Unlock()
}