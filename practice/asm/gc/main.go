package main

import (
	"fmt"
)

func main() {
	closure()
	closure2()

	// slice
	var s []int
	s = []int{1, 2}
	//go tool compile -d append main.go
	s = append(s, 3)
	//main.go:10:4: append: len-only update (in local slice)
	fmt.Println(s)
	//go tool compile -d typeassert main.go
	var a interface{}
	var b string
	a = "111"
	b = a.(string)
	println(b)
	//main.go:21:7: type assertion inlined
}

func closure() func() int {
	foo := 0
	// go tool compile -d closure main.go
	return func() int {
		foo++
		return foo
	}
	//main.go:7:9: heap closure, captured vars = foo
}

func closure2() func(int) int {
	sum := 0
	return func(x int) int {
		// println(x)
		sum += x
		println(sum)
		return sum
	}
	//main.go:39:9: heap closure, captured vars = sum
}
