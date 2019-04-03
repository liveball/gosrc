package main

import (
	"fmt"
	"unsafe"
)

func stringToBytes(s string) []byte {
	sp := *(*[2]uintptr)(unsafe.Pointer(&s))
	bp := [3]uintptr{sp[0], sp[1], sp[1]}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

func main() {
	fmt.Println(stringToBytes("abc"))
}
