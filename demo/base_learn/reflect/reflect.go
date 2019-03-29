package main

import (
	"fmt"
	"reflect"
)

//User for test.
type User struct {
	UserID int
	Name   string
}

func main() {
	test()
}

func test() {
	obj := User{UserID: 111, Name: "ddsdsds"}

	t := reflect.TypeOf(obj)

	newObj := reflect.New(t).Elem()

	for i := 0; i < t.NumField(); i++ {

		v := newObj.Field(i)

		switch t.Field(i).Type.Kind() {
		case reflect.String:
			v.SetString("阿呆")

		case reflect.Int:
			v.SetInt(123)

		}
	}

	fmt.Printf("obj(%+v)\n", obj)

	fmt.Printf("newObj(%+v)\n", newObj.Interface())

}

func dd() {
	x := User{UserID: 111}
	typ := reflect.TypeOf(x)
	// reflect.New返回的是*User 而不是User
	y := reflect.New(typ).Elem()
	for i := 0; i < typ.NumField(); i++ {
		// 根据每个struct field的type 设置其值
		fieldT := typ.Field(i)
		fieldV := y.Field(i)
		kind := fieldT.Type.Kind()
		if kind == reflect.Int {
			fieldV.SetInt(123)
		} else if kind == reflect.String {
			fieldV.SetString("wudaijun")
		}
	}

	fmt.Printf("y(%+v)\n", y.Interface())
}
