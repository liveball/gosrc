package main

import (
	"fmt"
	"sync"
)

func main() {
	// runner01b()

	a := make(chan int)
	a <- 1
}

func runner01b() {
	c := gen(2, 3, 5, 18, 31)
	out1 := sq(c)
	out2 := sq(sq(c))

	for i := range merge(out1, out2) {
		fmt.Println(i)
	}

}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
