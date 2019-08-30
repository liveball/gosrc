package main

import "fmt"

func main() {

	var c chan int
	c = make(chan int)
	go func() {
		c <- 1
	}()

	print(<-c)

	a := make([]int, 0)
	a = append(a, 1)
	fmt.Println(a)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		// println(x)
		sum += x
		println(sum)
		return sum
	}
}
