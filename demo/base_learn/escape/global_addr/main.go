package main

import (
	"fmt"
)

type obj struct{}

func main() {
	a := &obj{}
	fmt.Printf("%p\n", a)
	b := *a
	c := &b
	// fmt.Printf("%p\n", c)
	fmt.Println(a == c)
	fmt.Printf("%p\n", c)
}
