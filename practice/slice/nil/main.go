package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a []int
	// var b []int
	// b = make([]int, 0)

	_ = a
	// _ = b
	// fmt.Printf("slice a(%#v), slice b(%#v\n)",
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&a)),
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&b)),
	// 	// &b[0],//错误
	// )

	p := []byte{2, 3, 5}
	fmt.Println("append before ", p)
	p = appendByte(p, 7, 11, 13)
	fmt.Println("append after ", p)
	// p == []byte{2, 3, 5, 7, 11, 13}
}

// 结论：
// 1、描述一个不存在的切片的时候，就需要用到 nil 切片。比如函数在发生异常的时候，返回的切片就是 nil 切片，nil 切片的指针指向 nil。
// 2、空切片一般会用来表示一个空的集合。比如数据库查询，一条结果也没有查到，那么就可以返回一个空切片。
// 3、空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。
// 4、最后需要说明的一点是。不管是使用 nil 切片还是空切片，对其调用内置函数 append，len 和 cap 的效果都是一样的。

func appendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		fmt.Printf("before copy cap(%d) slice(%#v), newSlice(%#v)\n",
			(n+1)*2,
			(*reflect.SliceHeader)(unsafe.Pointer(&slice)),
			(*reflect.SliceHeader)(unsafe.Pointer(&newSlice)),
		)
		copy(newSlice, slice)
		slice = newSlice
		fmt.Printf("after copy slice(%#v), newSlice(%#v)\n",
			(*reflect.SliceHeader)(unsafe.Pointer(&slice)),
			(*reflect.SliceHeader)(unsafe.Pointer(&newSlice)),
		)
	}
	slice = slice[0:n]
	fmt.Printf("slice[0:n] (%+v)\n", slice)
	copy(slice[m:n], data)
	return slice
}
