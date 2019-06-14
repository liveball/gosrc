package main

import (
	"reflect"
	"unsafe"
)

func getg() interface{}

var g_goid_offset uintptr = func() uintptr {
	g := getg()
	if f, ok := reflect.TypeOf(g).FieldByName("goid"); ok {
		return f.Offset
	}
	panic("can not find g.goid field")
}()

func GetGoid() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + g_goid_offset))
	return *p
}

func main() {
	println(GetGoid())
}
