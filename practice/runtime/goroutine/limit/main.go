package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	taskCount = 100
	maxG      = 10
)

func fib(i int) {
	fmt.Printf("task id(%d)\n", i)
	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	limitCh := make(chan struct{}, maxG)

	statGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Printf("goroutine num(%d)\n", total)
		}
	}
	go statGroutine()

	for i := 0; i < taskCount; i++ {
		limitCh <- struct{}{}
		go func(i int) {
			fib(i)
			<-limitCh
		}(i)

	}

	select {}
}
