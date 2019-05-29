package main

import (
	"reflect"
)

//I def interface
type I interface {
	test1()
}

//T def struct
type T struct {
	A string
}

func (t *T) test1() {}

// func (t T) test1() {}

func main() {
	t := &T{}
	it := reflect.TypeOf((*I)(nil)).Elem()
	tv := reflect.TypeOf(t)
	println(tv.Implements(it))
}
