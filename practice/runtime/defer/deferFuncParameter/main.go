package main

import "fmt"

func main() {
	//deferFuncParameter()
	//
	//deferFuncParameter2()

	//officialDemo()

	//fmt.Println(foo())
	fmt.Println(foo2())

	//fmt.Println(deferFuncReturn())
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
	}()

	return i
}

func officialDemo() {
	i := 0

	defer fmt.Println(i)

	i++

	return
}

func foo() int {
	var i int

	defer func() {
		i++
	}()

	return 1
}

func foo2() int {
	var i int

	defer func() {
		i++
	}()

	return i
}
