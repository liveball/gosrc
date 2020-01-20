package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {
  var s=[]int{}

	fmt.Printf("s:%#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&s)),
	)
}

func test()  {
	slice:=[]int{11}

	//myMap:=make(map[int]*int)

	for _,v:=range slice {
		fmt.Println(&v)
	}

	//fmt.Println(myMap)
}