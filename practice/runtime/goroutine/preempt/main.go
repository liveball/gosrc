package main

import (
	"runtime"
	"sync/atomic"
)

func main() {
	// preemption()
	preemptionGC()
}

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

func preemption() {
	// Test that goroutines are preempted at function calls.
	N := 5
	c := make(chan bool)
	var x uint32
	for g := 0; g < 2; g++ {
		go func(g int) {
			for i := 0; i < N; i++ {
				for atomic.LoadUint32(&x) != uint32(g) {
					preempt()
				}
				// println("x", x)
				atomic.StoreUint32(&x, uint32(1-g))
			}
			c <- true
		}(g)
	}
	<-c
	<-c
}

func preemptionGC() {
	// Test that pending GC preempts running goroutines.
	P := 5
	N := 10
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
