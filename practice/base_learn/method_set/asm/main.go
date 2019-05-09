package main

import (
	"fmt"
	"unsafe"
)

type foo struct {
	a int
	b string
}

func (f foo) sub() {}

func main() {
	f := new(foo)
	*(*int)(unsafe.Pointer(f)) = 100
	fmt.Println(111, f)
	*(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(f)) + unsafe.Sizeof(f.b))) = "hello"
	fmt.Println(222, f)

	f.b = "world"
	fmt.Println(333, f)
}
