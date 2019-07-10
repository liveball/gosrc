package main

import "unsafe"

const ptrSize = 4 << (^uintptr(0) >> 63)

func main() {
	println(uintptr(0), ^uintptr(0), ^uintptr(0)>>63, 4<<(^uintptr(0)>>63), ^uint(0), 4<<(^uint(0)>>63))
	println(funcPC(test), funcPC(hello))
}

func test() {
	println("test")
}

func hello() {
	println("hello")
}

// funcPC returns the entry PC of the function f.
// It assumes that f is a func value. Otherwise the behavior is undefined.
// CAREFUL: In programs with plugins, funcPC can return different values
// for the same function (because there are actually multiple copies of
// the same function in the address space). To be safe, don't use the
// results of this function in any == expression. It is only safe to
// use the result as an address at which to start executing code.
//go:nosplit
func funcPC(f interface{}) uintptr {
	return **(**uintptr)(add(unsafe.Pointer(&f), ptrSize))
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}
