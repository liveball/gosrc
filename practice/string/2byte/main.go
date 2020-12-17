package main

import "fmt"

func main() {
	//slicestringcopy()

	t()
	//slicecopy()
}

//typedslicecopy
//slicestringcopy
//slicecopy
//memmove

func t() {
	var s = "a"
	for k, v := range s {
		fmt.Println(k,v)
	}

	//bs := *(*[]byte)(unsafe.Pointer(&s))

	bs := []byte(s)

	fmt.Println(string(bs))
}

//func slicestringcopy(to []byte, fm string) int {}

func slicestringcopy() {
	var a []byte
	a = make([]byte, 1)
	b := "bb"
	n := copy(a, b)
	fmt.Println(n, string(a))
}

//func slicecopy(to, fm slice, width uintptr) int

func slicecopy() {
	var a []int
	a = make([]int, 3)
	b := []int{1, 2, 3}
	n := copy(a, b)
	fmt.Println(n, a)
}

//func memmove(to, from unsafe.Pointer, n uintptr)
//TEXT runtime·memmove(SB), NOSPLIT, $0-24
