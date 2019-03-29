package main

import (
	"fmt"
)

func main() {
	// runtime.GOMAXPROCS(1)
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	ch := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		ch[i] = make(chan int)
		go func(c []chan int, s []int, j int) {
			fmt.Println("j", s[j])
			c[j] <- j
		}(ch, sli, i)
	}

	for i := range ch {
		fmt.Println("ch:", <-ch[i])
	}
}
