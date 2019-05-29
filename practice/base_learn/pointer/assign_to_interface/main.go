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

	// fmt.Printf("%#v\n",
	// 	(*reflect.StructField)(unsafe.Pointer(&p)),
	// )

	// var num float64 = 1.2345

	pointer := reflect.ValueOf(&p)
	value := reflect.ValueOf(p)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	convertPointer := pointer.Interface().(*Person)
	convertValue := value.Interface().(Person)

	println(&p, convertPointer)
	fmt.Println(convertValue == p, i == p)

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

	switch t.Kind() {
	case reflect.Struct:
		println(reflect.Struct)
		fmt.Println(reflect.Struct)
	case reflect.String:
		fmt.Println("string")
	}

	println("interface", i, newObj.Interface())
	fmt.Println("print newObj", newObj.Interface())

	fmt.Println("assert newObj.Interface() == p", newObj.Interface() == p)
	fmt.Println("assert newObj.Interface() == i", newObj.Interface() == i)
}
