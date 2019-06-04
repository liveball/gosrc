package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println(string2bytes("abc"))
	fmt.Println(bytes2string([]byte{65, 66, 67}))
}

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len}

	return *(*[]byte)(unsafe.Pointer(&bh))

}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}
