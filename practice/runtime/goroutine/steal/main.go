package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	ch := make(chan int)
	done := make(chan struct{})

	// runtime.GOMAXPROCS(1)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			// fmt.Println(idx, "go idx")
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send result")
			case <-done:
				fmt.Println(idx, "Exiting")
			}
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

	fmt.Println("Result: ", <-ch)
	close(done)

	time.Sleep(3 * time.Second)

}
