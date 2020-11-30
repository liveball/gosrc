package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		time.Sleep(time.Second)

		fmt.Println("goroutine 1 finished!")

		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 2)

		fmt.Println("goroutine 2 finished!")

		wg.Done()
	}()

	wg.Wait()

	fmt.Println("all goroutine finished")
}
