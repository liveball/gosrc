package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

type person struct {
	Age int `json:"age"`
}

func main() {
	psMap := make(map[int][]*person)

	for i := 6; i >= 0; i-- {
		psMap[1] = append(psMap[1],
			&person{
				Age: i,
			},
		)
	}

	fmt.Println(psMap)
	psNew, ok := psMap[1]
	if !ok || len(psNew) == 0 {
		return
	}

	fmt.Printf("psNew %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&psNew)),
	)

	println("排序前:")
	for _, v := range psNew {
		fmt.Println(v)
	}

	sort.Slice(psNew, func(i, j int) bool {
		return psNew[i].Age < psNew[j].Age
	})

	println("排序后:")
	for _, v := range psNew {
		fmt.Println(v)
	}
}
