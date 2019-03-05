package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type TestStructTobytes struct {
	data int64
}
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func main() {

	var testStruct = &TestStructTobytes{100}
	Len := unsafe.Sizeof(*testStruct)
	testBytes := &SliceMock{
		addr: uintptr(unsafe.Pointer(testStruct)),
		cap:  int(Len),
		len:  int(Len),
	}
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	fmt.Println("[]byte is : ", reflect.TypeOf(data), reflect.ValueOf(data))

	var ptestStruct = *(**TestStructTobytes)(unsafe.Pointer(&data))
	fmt.Println("ptestStruct.data is : ", reflect.TypeOf(ptestStruct), reflect.ValueOf(ptestStruct))
}
