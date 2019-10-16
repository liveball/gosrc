package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	taskCount = 10
	maxG      = 3

	limit = make(chan struct{}, maxG)
)

func fib(i int) {
	fmt.Printf("task id(%d)\n", i)
	//for {
	//	time.Sleep(1 * time.Second)
	//}
}

func main() {
	statGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Printf("goroutine num(%d)\n", total)
		}
	}
	go statGroutine()

	for i := 0; i < taskCount; i++ {
		go func(i int) {
			limit <- struct{}{}
			fib(i)
			<-limit
		}(i)

	}

	select {}
}

//This rule generalizes the previous rule to buffered channels.
//	It allows a counting semaphore to be modeled by a buffered channel:
//		the number of items in the channel corresponds to the number of active uses,
//        the capacity of the channel corresponds to the maximum number of simultaneous uses,
//        sending an item acquires the semaphore, and receiving an item releases the semaphore.
//	This is a common idiom for limiting concurrency.
//
//This program starts a goroutine for every entry in the work list,
//but the goroutines coordinate using the limit channel to ensure that at most three are running work functions at a time.

//var limit = make(chan int, 3)
//
//func main() {
//	for _, w := range work {
//		go func(w func()) {
//			limit <- 1
//			w()
//			<-limit
//		}(w)
//	}
//	select{}
//}
