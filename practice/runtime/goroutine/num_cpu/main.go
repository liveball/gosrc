package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
#define _GNU_SOURCE
#include <sched.h>
*/
// import "C"

var (
	a     string
	done  bool
	ncpus = runtime.NumCPU()
)

func main() {
	go setup()
	for !done {
	}
	print(a)

	var wg sync.WaitGroup
	wg.Add(ncpus)
	for i := 0; i < ncpus; i++ {
		j := i
		go func() {
			defer wg.Done()
			for {
				fmt.Printf("worker: %d\n", j)
				// fmt.Printf("worker: %d, CPU: %d\n", id, C.sched_getcpu())
			}
		}()
	}
	wg.Wait()
}

func setup() {
	a = "hello world"
	done = true
}
