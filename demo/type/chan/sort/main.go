package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 4)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	// go get(c)
	// go get(c)
	go get(c)
	time.Sleep(100 * time.Millisecond)
}

func get(cache <-chan int) {
	for {
		select {
		case i := <-cache:
			fmt.Println(i)
		}
	}
}
