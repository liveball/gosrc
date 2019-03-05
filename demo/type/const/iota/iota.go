package main

import "fmt"

const (
	A = iota
	B
	C = "c"
	D
	E = iota
	F
)

const (
	a, b = iota, iota << 10
	c, d
)

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(KB, MB, TB)
}
