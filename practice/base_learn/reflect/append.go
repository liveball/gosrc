package main

import (
	"fmt"
	"reflect"
)

func appendToSlice(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()
	value = reflect.Append(value, reflect.ValueOf(55))

	fmt.Println(value.Len()) // prints 1
}

func main() {
	arr := []int{}
	appendToSlice(&arr)
	fmt.Println(len(arr)) // prints 0
}

// reflect.Append works like append in that it returns a new slice value.

// You are assigning this value to the value variable in the appendToSlice function, which replaces the previous reflect.Value, but does not update the original argument.

// To make it more clear what's happening, take the equivalent function to your example without reflection:

func appendToSlice2(arrPtr *[]int) {
	value := *arrPtr
	value = append(value, 55)
	fmt.Println(len(value))
}

// What you need to use is the Value.Set method to update the original value:

func appendToSlice3(arrPtr interface{}) {
	valuePtr := reflect.ValueOf(arrPtr)
	value := valuePtr.Elem()

	value.Set(reflect.Append(value, reflect.ValueOf(55)))

	fmt.Println(value.Len())
}
