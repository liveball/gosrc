package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int{
	return a+b
}

func main(){
	author := "呵呵呵"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	i := 1
	//vv := reflect.ValueOf(i)
	//vv.SetInt(10) //panic: reflect: reflect.Value.SetInt using unaddressable value
	//fmt.Println(i)

	//修改
	vvv := reflect.ValueOf(&i)
	vvv.Elem().SetInt(10)
	fmt.Println(i)

	v:=reflect.ValueOf(Add)

	if v.Kind()!=reflect.Func{
		return
	}


	t:=v.Type()
	argv:=make([]reflect.Value, t.NumIn())

	method,ok:=t.MethodByName("Add")
	fmt.Println(111, method, ok, t.NumMethod(), t.NumIn())
	for i:=range argv{
		if t.In(i).Kind()!=reflect.Int{
			return
		}
		fmt.Println(222, i)
		argv[i]=reflect.ValueOf(i)
	}
	res:=v.Call(argv)
	if len(res)!=1|| res[0].Kind()!=reflect.Int{
		return
	}

	fmt.Println(res[0].Int())
}
