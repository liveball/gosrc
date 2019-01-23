package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// a := []byte{1, 2, 3, 3}
	// b := append(a[:2], 4, 4)
	// fmt.Printf("a(%+v) len(%d) cap(%d) %#v\n",
	// 	a, len(a), cap(a),
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&a)),
	// )
	// fmt.Printf("b(%+v) len(%d) cap(%d) %#v\n",
	// 	b, len(b), cap(b),
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&b)),
	// )

	c := []byte{1, 2, 3, 3}
	d := append(c[:2], 4, 4, 4)
	fmt.Printf("c(%+v) len(%d) cap(%d) %#v\n",
		c, len(c), cap(c),
		(*reflect.SliceHeader)(unsafe.Pointer(&c)),
	)
	fmt.Printf("d(%+v) len(%d) cap(%d) %#v\n",
		d, len(d), cap(d),
		(*reflect.SliceHeader)(unsafe.Pointer(&d)),
	)
}
