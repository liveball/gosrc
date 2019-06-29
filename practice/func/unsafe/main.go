package main

import "unsafe"

type person struct {
	age int

	name unsafe.Pointer
}

func main() {
	_p := &person{
		age: 10,
	}
	fn := *(*func(*person, unsafe.Pointer) bool)(unsafe.Pointer(&_p.name))

	p := &person{
		age: 20,
	}
	ok := fn(p, _p.name)

	println(ok)
}
