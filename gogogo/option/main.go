package main

import "fmt"

type Foo struct {
	key    string
	option Option

	// ...
}

type Option struct {
	num int
	str string
}

type ModOption func(option *Option)

func New(key string, modOption ModOption) *Foo {
	option := Option{
		num: 100,
		str: "hello",
	}

	modOption(&option)

	return &Foo{
		key:    key,
		option: option,
	}
}

// ...
func main() {
	modOp1 := func(option *Option) {
		// 调用方只设置 num，覆盖默认值
		option.num = 200
	}
	// modOp2 := func(option *Option) {
	// 	// 调用方只设置 num，覆盖默认值
	// 	option.num = 300
	// }

	f := New("iamkey", modOp1)

	fmt.Println("num:", f.option.num)
}
