package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var arr [8]int = [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("arr %#v\n",
		unsafe.Pointer(&arr),
	)

	fmt.Printf("arr[0] %#v\n",
		unsafe.Pointer(&arr[0]),
	)

	fmt.Printf("arr[1] %#v\n",
		unsafe.Pointer(&arr[1]),
	)

	var s1 []int
	s1 = arr[1:4]
	c1 := cap(s1)
	fmt.Println("s1_len", len(s1))
	fmt.Println("s1_cap", c1)

	fmt.Printf("s1 %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&s1)),
	)

	var s2 = arr[2:4]
	fmt.Println("s2_len", len(s2))
	fmt.Println("s2_cap", cap(s2))

	fmt.Printf("s2 %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&s2)),
	)
}
