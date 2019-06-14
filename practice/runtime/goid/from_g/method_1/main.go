package main

import (
	"runtime"
	"strings"
	"unsafe"
)

var offsetDictMap = map[string]int64{
	"go1.10": 152,
	"go1.9":  152,
	"go1.8":  192,
}

var g_goid_offset = func() int64 {
	goversion := runtime.Version()
	for key, off := range offsetDictMap {
		if goversion == key || strings.HasPrefix(goversion, key) {
			return off
		}
	}
	panic("unsupport go verion:" + goversion)
}()

func getg() unsafe.Pointer

func GetGoid() int64 {
	g := getg()
	p := (*int64)(unsafe.Pointer(uintptr(g) + g_goid_offset))
	return *p
}

func main() {
	println(GetGoid())
}
