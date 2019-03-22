package main

import "fmt"

type foo struct {
}

func main() {
	judge(foo{})
	// judge(&foo{})
}

func judge(i interface{}) {
	switch i.(type) {
	case struct{}:
		fmt.Println("struct")
	case *struct{}:
		fmt.Println("*struct")
	default:
		fmt.Println("no")
	}
}

// 传入一个结构体，无论是用struct{}还是&struct{}，都得到no，这是为啥
