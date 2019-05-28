package main

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

type person struct {
	name string
	age  int
}

func main() {
	p := &person{}
	// fmt.Printf("before modify p: %+v\n", p)
	spew.Dump(111, p)
	modifyPerson(p) //p 是实参
	spew.Dump(333, p)

	// fmt.Printf("after modify p: %+v\n", p)
}

func modifyPerson(p *person) { //p 是形参
	// 1:error, p 是指针变量，本身存储地址，函数内部被修改，调用完成释放

	// p = &person{
	// 	name: "xiaohan",
	// }

	// 2/3:不修改指针的值，只修改p指向的内容
	*p = person{
		name: "xiaohan",
	}

	// p.name = "xiaohan"

	spew.Dump(222, p)

	fmt.Printf("%#v\n",
		(*reflect.StructField)(unsafe.Pointer(&p)),
	)

	//结论：
	//修改了形参的值，即指针变量的值，函数调用完成，这个临时赋值的地址被释放，实参指针指向的内容并没有发生改变
}
