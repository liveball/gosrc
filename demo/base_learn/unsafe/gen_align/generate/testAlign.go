package main

import (
	"fmt"
	"reflect"
)

var v = struct {
	a complex128
	b uintptr
	c bool
	d uint16
	e uint8
}{0 + 0i, 0, false, 0, 0}

func main() {
	fmt.Printf("%#T\n", v)
	t := reflect.TypeOf(v)
	fmt.Printf("结构体大小：%v\n", t.Size())
	for i := 0; i < t.NumField(); i++ {
		showAlign(t, i)
	}
}

func showAlign(v reflect.Type, i int) {
	sf := v.Field(i)
	fmt.Printf("字段 %10v，大小：%2v，对齐：%2v，字段对齐：%2v，偏移：%2v\n",
		sf.Type.Kind(),
		sf.Type.Size(),
		sf.Type.Align(),
		sf.Type.FieldAlign(),
		sf.Offset,
	)
}
