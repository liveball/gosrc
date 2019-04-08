package main

import (
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
		go func(id int) {
			defer wg.Done()
			for {
				// fmt.Printf("worker: %d, CPU: %d\n", id)
				// fmt.Printf("worker: %d, CPU: %d\n", id, C.sched_getcpu())
			}
		}(i)
	}
	wg.Wait()
}

func setup() {
	a = "hello world"
	done = true
}
