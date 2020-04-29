package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []byte("")
	fmt.Println("s=>", len(s), cap(s), (*reflect.SliceHeader)(unsafe.Pointer(&s)))

	s1 := append(s, 'a')

	fmt.Println("s1=>", len(s1), cap(s1), (*reflect.SliceHeader)(unsafe.Pointer(&s1)))

	s2 := append(s, 'b')

	fmt.Println("s2=>", len(s2), cap(s2), (*reflect.SliceHeader)(unsafe.Pointer(&s2)))

	//fmt.Println(s1, "==========", s2)
	fmt.Println(string(s1), "==========", string(s2))
}

// 出现个让我理解不了的现象, 注释时候输出是 b ========== b
// 取消注释输出是 [97] ========== [98] a ========== b
