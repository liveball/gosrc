package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	default:
		fmt.Println("default")
	}

	s := [3]int{1, 2, 3}
	a := s[1:2:cap(s)]
	b := s[:0]
	fmt.Println(cap(s), a, b)
}
