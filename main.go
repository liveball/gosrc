package main

import "fmt"

func main() {
	println("hello")

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
