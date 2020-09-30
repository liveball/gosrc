package main

import (
	"fmt"
	"reflect"
)

type Bar struct {
	I int
}

type Foo struct {
	J int
	//K map[string]struct{} //包含map、slice 等不能比较的类型
}

func main() {
	var a = Bar{}
	var b = Foo{}

	//fmt.Println(a == b)
	fmt.Println(reflect.DeepEqual(a, b))

	a1 := Foo(b)

	//invalid operation: a1 == b (struct containing map[string]struct {} cannot be compared)
	fmt.Println(a1 == b)
	fmt.Println(reflect.DeepEqual(a1, b))
}
