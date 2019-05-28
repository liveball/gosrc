package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//Person for person info
type Person struct {
	//Name for name
	Name string
}

func main() {
	var p Person
	p.Name = "小明"

	var i interface{}
	i = p

	fmt.Println("assert p == i", p == i, *(*Person)(unsafe.Pointer(&p)), i)

	if ii, ok := i.(Person); ok {
		ii.Name = "小黑"
		// fmt.Println(ii)
		fmt.Println("assert p == ii", p == ii)
	}

	t := reflect.TypeOf(i)
	newObj := reflect.New(t).Elem()
	for i := 0; i < t.NumField(); i++ {
		v := newObj.Field(i)
		switch t.Field(i).Type.Kind() {
		case reflect.String:
			v.SetString("阿呆")
		}
	}

	println("interface", i, newObj.Interface())
	fmt.Println("print newObj", newObj.Interface())

	fmt.Println("assert newObj.Interface() == p", newObj.Interface() == p)
	fmt.Println("assert newObj.Interface() == i", newObj.Interface() == i)
}
