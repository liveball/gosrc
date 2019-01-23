package main

import (
	"fmt"
	"reflect"
)

func test() {
	var a interface{}
	a = []string{"1", "2"}

	fmt.Println(reflect.TypeOf(a).Comparable())
}
