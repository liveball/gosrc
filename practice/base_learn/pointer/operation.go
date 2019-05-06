package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := struct {
		s string
		x int
	}{"abc", 100}
	fmt.Printf("before:%#v\n", a)

	p := uintptr(unsafe.Pointer(&a))
	p += unsafe.Offsetof(a.x)
	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 200
	fmt.Printf("after:%#v\n", a)
}
