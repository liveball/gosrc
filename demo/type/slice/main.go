package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 10
		// fmt.Println(arr) // [10 2 3]
	}(x)
	// fmt.Println(x) // [1 2 3]

	func(arr *[3]int) { //传数组指针改变数组元素的值
		(*arr)[0] = 10
		// fmt.Println(*arr) // [10 2 3]
	}(&x)
	// fmt.Println(x) //[10 2 3]

	y := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 100
		// fmt.Println(x) // [100 2 3]
	}(y)
	// fmt.Println(y) //[100 2 3]

	table := make([][]int, 2)
	for i := range table {
		table[i] = make([]int, 4)
	}
	// fmt.Printf("%#v, %#v\n",
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&table)),
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&table)).Data,
	// )

	twoDimension() //创建二维数组

	forrange() //循环值拷贝

	sliceCopyToSlice()
	stringCopyToByteSlice()

	//Go 数组是值类型， 赋值和函数传参操作都会复制整个数组数据， 下面三个数组的内存地址都不同。
	arrayA := [2]int{100, 200}
	var arrayB [2]int
	arrayB = arrayA
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)

	testArray(arrayA)

	slicePointer()
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

func changeSlice() {
	a := []int{1, 2, 3}
	b := append(a, 10) //a/b依赖不同的底层数组

	fmt.Printf("%#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&a)),
		(*reflect.SliceHeader)(unsafe.Pointer(&b)),
	)
}

func sliceInfo(x []int) {
	fmt.Printf("len is %d ,cap is %d,  slice is %v address %p\n", len(x), cap(x), x, &x)
}

func forrange() {
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("value = %d , value-addr = %x , slice-addr = %x\n", value, &value, &slice[index])
	}
}

func sliceCopyToSlice() {
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n, slice)
}

func stringCopyToByteSlice() {
	slice := make([]byte, 3)
	n := copy(slice, "abcdef")
	fmt.Println(n, slice)
}

func twoDimension() {
	h, w := 2, 4
	raw := make([]int, h*w)

	for i := range raw {
		raw[i] = i
	}

	// 初始化原始 slice
	fmt.Println(raw, &raw[4]) // [0 1 2 3 4 5 6 7] 0xc420012120

	table := make([][]int, h)
	for i := range table {
		// 等间距切割原始 slice，创建动态多维数组 table
		// 0: raw[0*4: 0*4 + 4]
		// 1: raw[1*4: 1*4 + 4]
		table[i] = raw[i*w : i*w+w]
		fmt.Println(1111, table)
	}

	fmt.Println(table, &table[1][0]) // [[0 1 2 3] [4 5 6 7]] 0xc420012120

	fmt.Printf("%#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&raw)),
		(*reflect.SliceHeader)(unsafe.Pointer(&table)),
	)
}

func slicePointer() {
	var a = [...]int{1, 2}
	b := a[:0]
	b = append(b, 3) //改变数组第一个元素的值
	// fmt.Println(a, b)
	// fmt.Printf("%p, %p\n", a, b)
	fmt.Printf("%#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&a)),
		(*reflect.SliceHeader)(unsafe.Pointer(&b)),
	)

	a2 := []int{1, 2}
	b2 := append(a2, 3)
	fmt.Printf("%#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&a2)),
		(*reflect.SliceHeader)(unsafe.Pointer(&b2)),
	)
	// &reflect.SliceHeader{Data:0xc420016140, Len:2, Cap:2}, &reflect.SliceHeader{Data:0xc42001c180, Len:3, Cap:4}

	c := make([]int, 2, 10)

	d := c[:0]
	d = append(c, 3)

	fmt.Printf("%#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&c)).Data, //弱引用，对象可被垃圾回收，不可达
		(*reflect.SliceHeader)(unsafe.Pointer(&d)).Data,
	)

	type sliceStruct struct {
		array uintptr //unsafe.Pointer 强引用
		len   int
		cap   int
	}

	fmt.Printf("%#v, %#v\n",
		(*sliceStruct)(unsafe.Pointer(&c)).array, //强引用，使用垃圾回收，对象不可达
		(*sliceStruct)(unsafe.Pointer(&d)).array,
	)

	// 从 Go 的内存地址中构造一个 slice
	var ptr uintptr
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{ptr, 1, 1}
	s := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Printf("s %#v\n", s)

	// 如果想从 slice 中得到一块内存地址，可以这样做：
	s2 := make([]byte, 200)
	ptr2 := unsafe.Pointer(&s2[0])
	fmt.Printf("ptr2 %#v\n", ptr2)

}

// var c [100 << 20]byte
// runtime.SetFinalizer(&c, func(o interface{}) {
// 	println("c dead")
// })
// p := &c
// p := uintptr(unsafe.Pointer(&c)) //uintptr不可达，c dead
// _ = p
// runtime.GC()
// runtime.KeepAlive(p) //确保p指针活着,这样gc完成后，可以观察p指针是否可以让c保持可达状态
