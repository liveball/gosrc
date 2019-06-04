package main

import (
	"runtime"
	"sync/atomic"
	"testing"
)

// The function is used to test preemption at split stack checks.
// Declaring a var avoids inlining at the call site.
var preempt = func() int {
	var a [128]int
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func TestPreemptionGC(t *testing.T) {
	if runtime.GOARCH == "wasm" {
		t.Skip("no preemption on wasm yet")
	}

	// Test that pending GC preempts running goroutines.
	P := 5
	N := 10
	if testing.Short() {
		P = 3
		N = 2
	}
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(P + 1))
	var stop uint32
	for i := 0; i < P; i++ {
		go func() {
			for atomic.LoadUint32(&stop) == 0 {
				preempt()
			}
		}()
	}
	for i := 0; i < N; i++ {
		runtime.Gosched()
		runtime.GC()
	}
	atomic.StoreUint32(&stop, 1)
}
