package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf([]int{1, 2}) // a reflect.Value
	fmt.Println(v)                    // "3"
	fmt.Println(v.String())           // NOTE: "<int Value>"
}
