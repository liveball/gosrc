package main

import "fmt"

func main() {
	var a, b byte
	a = 255
	b = 1

	fmt.Printf("a:%d, %08b\n", a, a)
	fmt.Printf("b:%d, %08b\n", b, b)

	var c uint16
	c = 256
	fmt.Printf("c:%d, %08b\n", c, c)

}
