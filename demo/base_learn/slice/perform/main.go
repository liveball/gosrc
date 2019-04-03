package main

import "fmt"

func main() {
	a := make([]int, 0, 1)
	a = a[:0] //清空 1
	fmt.Println(a, len(a), cap(a))

	// fmt.Println(a, len(a), cap(a))
	// a = make([]int, 0, 1) //清空 2
	// fmt.Println(a, len(a), cap(a))

}
