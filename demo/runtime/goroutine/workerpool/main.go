package main

import (
	"fmt"
	"runtime"
	"time"
)

func handle(i interface{}) {
	j, ok := i.(int)
	if !ok {
		return
	}

	fib(j)
}

func fib(i int) {
	fmt.Printf("g(%d)\n", i)
	for {
		// fmt.Printf("data (%d)\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	maxJob := 10
	// maxWorker := 10
	// dispatcher := NewDispatcher(maxWorker)
	// dispatcher.Run()
	// JobQueue = make(chan Job, maxJob)
	for i := 0; i < maxJob; i++ {
		// JobQueue <- Job{
		// 	Data: i,
		// 	Proc: handle,
		// }

		go func(i int) {
			fib(i)
		}(i)
	}

	statGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Printf("goroutine num(%d)\n", total)
		}
	}
	go statGroutine()
	select {}
}
