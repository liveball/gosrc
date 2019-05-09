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
	var a interface{}
	var b interface{} = (*interface{})(nil)

	println(a, a == nil)
	println(b, b == nil)

	fmt.Printf("a type: %#v\tvalue: %#v \n",
		reflect.TypeOf(a),
		reflect.ValueOf(a),
	)

	var (
		v  interface{}
		r  io.Reader
		f  *os.File
		fn os.File
	)

	fmt.Println("interface{}", v == nil)
	fmt.Println("io.Reader", r == nil)
	fmt.Println("*os.File", f == nil)

	v = r
	fmt.Println("v interface{}=r io.Reader", v == nil)

	v = f
	fmt.Println("v interface{}=f *os.File", v == nil)
	r = f
	fmt.Println("r io.Reader=f *os.File", r == nil)

	v = fn
	fmt.Println("v interface{}=fn os.File", v == nil)

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
