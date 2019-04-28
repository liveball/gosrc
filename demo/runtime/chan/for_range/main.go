package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			ch <- j
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(ch)

	for i := range ch {
		fmt.Println("ch:", i)
	}
}
