package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	gcflags := []string{"-+", "-N", "-m", "-l"}
	for i := 0; i < len(gcflags); i++ {
		if gcflags[i] == "-N" { //i=1
			copy(gcflags[i:], gcflags[i+1:])
			gcflags = gcflags[:len(gcflags)-1] //gcflags [-+ -m -l -l]
			i--
		}
	}
	fmt.Println(gcflags)
}

func del() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("remove before a %#v \n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))

	removeByRef(3, a)
	// removeByPointer(3, &a)
}

func removeByRef(m int, s []int) {
	tt := make([]int, 10, 10)
	for i := 0; i < len(s); i++ {
		if m == s[i] {

			fmt.Println("1", s[:i])
			fmt.Println("2", s[i+1:])

			t := append(s[:i], s[i+1:]...) //依赖同一个底层数组

			//copy t
			count := copy(tt, t)
			fmt.Printf("tt val(%+v) tt ref(%p) count(%d)\n", tt, tt, count)

			// r := (*reflect.SliceHeader)(unsafe.Pointer(&s))
			// r.Data = uintptr((*reflect.SliceHeader)(unsafe.Pointer(&t)).Data)

			fmt.Printf("s val(%+v) element(%d) ref(%p) addr(%p) slice %#v \n",
				s, s[3], s, &s[3],
				(*reflect.SliceHeader)(unsafe.Pointer(&s)),
			)

			s = t //s 和 t同为slice，引用的底层数组的地址一样，所以此赋值操作无效

			fmt.Printf("t val(%+v) element(%d) ref(%p) addr(%p) slice %#v \n",
				t, t[3], t, &t[3],
				(*reflect.SliceHeader)(unsafe.Pointer(&t)),
			)

		}
	}
}

func removeByPointer(m int, s *[]int) {
	for i := 0; i < len(*s); i++ {
		if m == (*s)[i] {
			t := append((*s)[:i], (*s)[i+1:]...) //依赖同一个底层数组

			*s = t
			fmt.Printf("*s (%p) \n", *s)

			fmt.Printf("t val(%v) ref(%p) addr(%p)  slice %#v \n",
				t, t, &t,
				(*reflect.SliceHeader)(unsafe.Pointer(&t)),
			)
		}
	}
}
