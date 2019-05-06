package main

import "fmt"

func main() {
	a := struct{}{}
	b := struct{}{}
	println("a=", &a)
	println("b=", &b)

	fmt.Printf("a:%p\n", &a)
	fmt.Printf("b:%p\n", &b)
}
