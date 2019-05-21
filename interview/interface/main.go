package main

import (
	"fmt"
	"reflect"
)

//People def people
type People interface {
	Show()
}

//Student def stu instance
type Student struct{ Name string }

//Show show stu info
func (stu *Student) Show() { println(stu.Name) }

func live() People {
	var stu *Student
	return stu
}

func main() {
	a := live()
	fmt.Printf("type: %#v\nvalue: %#v \n",
		reflect.TypeOf(a),
		reflect.ValueOf(a),
	)
	if a == nil {
		println("nil")
	} else {
		println("not nil")
	}
}
