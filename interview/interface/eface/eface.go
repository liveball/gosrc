package main

import (
	"fmt"
	"unsafe"
)

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type tflag uint8
type nameOff int32
type typeOff int32
type textOff int32

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}

func main() {
	var ti interface{}
	var a int = 100
	ti = a

	fmt.Println("type:", *(*eface)(unsafe.Pointer(&ti))._type)
	fmt.Println("data:", *(*int)((*eface)(unsafe.Pointer(&ti)).data))
	fmt.Println((*eface)(unsafe.Pointer(&ti)))
}
