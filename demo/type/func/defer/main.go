package main

type domain struct {
	do string
}

func main() {
	a := domain{
		do: "aa",
	}

	// fmt.Printf("1 a val(%+v) addr(%p)\n", a, &a)

	// defer fd(a) //{bb} {aa} // defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体。

	defer func() {
		// fmt.Printf("defer  a val(%+v) addr(%p)\n", a, &a)
		fd(a) //a 为引用传递
	}() //{bb} {bb}

	// defer func(a domain) {
	// 	fmt.Printf("defer  a val(%+v) addr(%p)\n", a, &a)
	// 	fd(a) //a 为值传递
	// }(a) //{bb} {bb}

	a = domain{
		do: "bb",
	}

	// fmt.Printf("2 a val(%+v) addr(%p)\n", a, &a)
}

func fd(a domain) {
	// fmt.Printf("fd a val(%+v) addr(%p)\n", a, &a)
}
