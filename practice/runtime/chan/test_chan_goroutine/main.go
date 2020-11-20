package main

import (
	"fmt"
	"sync"
)

func main() {
	w := make(chan int, 5)
	// t := make(chan bool)
	var i int

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for s := 0; s < 10; s++ {
			w <- s
		}
		// t <- true
		close(w)
		wg.Done()
	}()

	go func() {
		for i = range w {
			fmt.Println(i)
		}
		wg.Done()
	}()

	// <-t
	wg.Wait()
	fmt.Println("end")
}
