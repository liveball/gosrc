package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	base()
	arr := []int{1, 2, 3}
	printArr(arr)

	var a []interface{}
	a = make([]interface{}, 0, 2)
	a = append(a, 1)
	a = append(a, 2)
	fmt.Println(a)
}

func base() {
	var x int = 100

	var a interface{} = x
	var b interface{} = &x

	fmt.Printf("unsafe.Sizeof of a(%d) b(%d) reflect.ValueOf a(%#v) b(%#v)  \n",
		unsafe.Sizeof(a), unsafe.Sizeof(b), reflect.ValueOf(a), reflect.ValueOf(b))

	y := interface{}(x).(int)
	z := interface{}(x)

	fmt.Printf("x(%#v) y(%#v) z(%#v) \n",
		reflect.ValueOf(x), reflect.ValueOf(y), reflect.ValueOf(z))
}

func printArr(arr interface{}) {
	a, ok := arr.([]int)
	if ok {
		for _, v := range a {
			fmt.Printf("v(%v)\n", v)
		}
	}

	slice := reflect.ValueOf(arr)

	if slice.Kind() != reflect.Slice {
		panic(fmt.Sprintf("%+v is not a slice", slice))
	}
	fmt.Printf("slice.Slice(0, 0).Interface():(%+v)\n", slice.Slice(0, 0).Interface())
}
