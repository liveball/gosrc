package main

import (
	"fmt"
	"go/types"
	"reflect"
	"unsafe"
)

type b byte

const (
	k b = iota
	m
	n
)

func main() {
	// testIsAlias()
	var a int
	_ = a

	types.DefPredeclaredTestFuncs()

	fmt.Println("m:", m, "n:", n)
	//iota
	fmt.Printf("unsafe.Sizeof of k(%d) m(%d)  reflect.ValueOf k(%#v) m(%#v) \n",
		unsafe.Sizeof(k), unsafe.Sizeof(m), reflect.ValueOf(k), reflect.ValueOf(m),
	)
}
