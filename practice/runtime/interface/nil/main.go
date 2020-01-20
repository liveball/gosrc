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

	fmt.Println(a, a == nil)
	fmt.Println(b, b == nil)

	fmt.Printf("a type: %#v\t value: %#v \n",
		reflect.TypeOf(a),
		reflect.ValueOf(a).String(),
	)

	fmt.Printf("b type: %#v\t value: %#v \n",
		reflect.TypeOf(b).String(),
		reflect.ValueOf(b).String(), //*(*[2]uintptr)(unsafe.Pointer(&b))
	)

	fmt.Println("=======================")

	type foo struct{}

	var f foo
	var f2 *foo
	a = f
	b1 := f2

	fmt.Println(a, a == nil)
	fmt.Println(b1, b1 == nil)

	fmt.Printf("a1 type: %#v\t value: %#v \n",
		reflect.TypeOf(a).String(),
		reflect.ValueOf(a).String(),
	)

	fmt.Printf("b1 type: %#v\t value: %#v \n",
		reflect.TypeOf(b1).String(),
		reflect.ValueOf(b1).String(),
	)

	fmt.Println(IsNilPtr(a), IsNilPtr(b1))

	fmt.Println("=======================")

	var sl []string
	yes := sl == nil
	fmt.Println(yes)

	fmt.Printf("sl type: %#v\t value: %#v \n",
		reflect.TypeOf(sl).String(),
		reflect.ValueOf(sl).String(),
	)
	//test()
}

func IsNilPtr(x interface{}) bool {
	v := reflect.ValueOf(x)
	return v.Kind() == reflect.Ptr && v.IsNil()
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

func test() {

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
