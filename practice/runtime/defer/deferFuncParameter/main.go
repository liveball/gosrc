package main

import "fmt"

func main() {
	//deferFuncParameter()
	//
	//deferFuncParameter2()

	//officialDemo()

	//fmt.Println(deferFuncReturn())

	//fmt.Println(foo())

	fmt.Println(foo2())
}

func deferFuncParameter() {
	var a = 1

	defer fmt.Println(a)

	a = 2

	return
}

func deferFuncParameter2() {
	var arr = [3]int{1, 2, 3}

	defer printArr(&arr)

	arr[0] = 10

	return
}

func printArr(arr *[3]int) {
	for k, v := range arr {
		fmt.Println(k, arr[k], v)
	}
}

func deferFuncReturn() (res int) {
	i := 1

	defer func() {
		res++
		fmt.Println("res:", i)
	}()

	return i
}

func officialDemo() {
	i := 0

	defer fmt.Println(i)

	i++

	return
}

func foo() int { //defer 无法操作返回值
	var i int
	i = 2

	defer func() {
		i++
		fmt.Println("i:", i)
	}()

	return 1
}

func foo2() int { //defer 无法操作返回值
	var i int

	defer func() {
		i++
	}()

	return i
}
