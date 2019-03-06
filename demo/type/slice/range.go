package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	println("=====array=========")
	a := [3]int{0, 1, 2}
	println(&a, &a[0])
	for k, v := range a {
		if k == 0 {
			a[1], a[2] = 998, 999
			fmt.Println(a)
		}

		a[k] = v + 100
		vv := v
		println(&a[k], &v, vv, &vv)
	}
	fmt.Println(a)

	println("=====slice=========")
	s := []int{1, 2, 3, 4, 5}
	println(&s, &s[0])
	fmt.Printf("s before:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 6
			println(i, &s[i])
			fmt.Printf("s after:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)))
		}
		vv := v
		fmt.Println(i, v, &v, vv, &vv)
	}
	fmt.Println(s)
}
