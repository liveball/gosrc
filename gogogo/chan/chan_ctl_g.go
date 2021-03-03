package main

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	time.Sleep(time.Second)

	ch <- 1
}

func main() {
	channels := make([]chan int, 2)

	for i := 0; i < 2; i++ {
		channels[i] = make(chan int)
		go Process(channels[i])
	}

	for i, ch := range channels {
		// <-ch
		// fmt.Println("routine ", i, " quit")

		select {
		case elem, ok := <-ch:
			if !ok {
				return
			}

			fmt.Println("routine ", i, "get elem:", elem)
		}
	}

}
