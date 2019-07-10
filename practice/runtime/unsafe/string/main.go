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

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))                                                    // case 1
	hdr.Data = uintptr(unsafe.Pointer((*string)(unsafe.Pointer(reflect.ValueOf(new(string)).Pointer())))) // case 6 (this case)
	hdr.Len = len(s)

	fmt.Println(hdr)

	s2 := *(*string)(unsafe.Pointer(&hdr))
	println("modify s after:", "len=", len(s2), "s2=")

	// StringHeader.Data 是uintptr 类似，为弱引用，可能被垃圾回收器回收掉
	// println(s2) //panic fatal error: unexpected signal during runtime execution

	// var n name
	// b := (*[4]byte)(unsafe.Pointer(n.bytes))
	// hdr.Data = unsafe.Pointer(&b[3])
	// hdr.Len = int(b[1])<<8 | int(b[2])

	b := []byte("hello")
	fmt.Printf("切片第一个元素地址: %#v\n", &b[0])

	strNoAlloc := byteToStringNoAlloc(b)
	shNoAlloc := (*reflect.StringHeader)(unsafe.Pointer(&strNoAlloc))
	fmt.Printf("不分配内存的底层数据: %#v\n", shNoAlloc)

	str := byteToString(b)
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("分配内存的底层数据: %#v\n", sh)

}

func byteToString(b []byte) string {
	return string(b)
}

func byteToStringNoAlloc(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	sh := reflect.StringHeader{Data: uintptr(unsafe.Pointer(&b[0])), Len: len(b)}
	return *(*string)(unsafe.Pointer(&sh))
}
