package main

import (
	"fmt"
	"go/types"
	"io"
	"os"
	"reflect"
)

//使用 -gcflags "-N -l" 参数关闭编译器代码优化
//go build -gcflags "-N -l" -o main main.go

func main() {
	var a interface{} = nil
	var b interface{} = (*interface{})(nil)

	println(a, a == nil)
	println(b, b == nil)

	var (
		v  interface{}
		r  io.Reader
		f  *os.File
		fn os.File
	)

	fmt.Println(v == nil)
	fmt.Println(r == nil)
	fmt.Println(f == nil)
	v = r
	fmt.Println(v == nil)
	v = fn
	fmt.Println(v == nil)
	v = f
	fmt.Println(v == nil)
	r = f
	fmt.Println(r == nil)

	// fmt.Printf("a type: %#v \n\n value:  %#v \n",
	// 	reflect.TypeOf(a),
	// 	reflect.ValueOf(a),
	// )
}

func dumpObj() {
	for _, name := range types.Universe.Names() {
		// println(name)
		if obj, _ := types.Universe.Lookup(name).(*types.TypeName); obj != nil {
			fmt.Printf("obj(%#v) \n", reflect.ValueOf(obj))
			println()
		}
	}
}
