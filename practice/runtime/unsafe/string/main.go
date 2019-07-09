package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type name struct {
	bytes *byte
}

func main() {
	var s string
	s = "hello world"
	println("modify s before:", "len=", len(s), "s=", s)

	p := (*string)(unsafe.Pointer(reflect.ValueOf(new(string)).Pointer()))
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
	hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
	hdr.Len = len(s)

	fmt.Println(hdr)

	s2 := *(*string)(unsafe.Pointer(&hdr))
	println("modify s after:", "len=", len(s2), "s2=")

	// println(s2)//panic fatal error: unexpected signal during runtime execution

	// var n name
	// b := (*[4]byte)(unsafe.Pointer(n.bytes))
	// hdr.Data = unsafe.Pointer(&b[3])
	// hdr.Len = int(b[1])<<8 | int(b[2])
}
