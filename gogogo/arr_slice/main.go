package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var a = [2]int{0, 1}
	fmt.Println("before modifyArr:", a)
	modifyArr(a)
	fmt.Println("after modifyArr:", a)

	var b = []int{0, 1}
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))
	fmt.Println("before modifySlice:", b)
	modifySlice(b)
	fmt.Println("after modifySlice:", b)
}

func modifyArr(a [2]int) {
	a[0] = 10
	a[1] = 11

	fmt.Printf("%#v\n", a)
}

// 切片的扩容
// 当在调用append()对切片进行追加时，如果添加元素的个数加上原有切片长度大于原有容量的话就会触发扩容操作，扩容操作时，
// 会根据以下的几个条件进行对扩容大小的选择：

// 1.当期望容量（即当前元素长度加上要添加的元素个数）大于当前容量的两倍时，按照期望容量的大小作为目标容量大小。
// 2.如果条件1不满足，并且当前切片容量（1.16版本之前是长度）小于1024时）目标容量大小为当前容量大小的两倍。
// 3.如果条件1，2都不满足则循环计算并增加目标容量，每次循环增加原有容量的25%，直至目标容量超过期望容量。
func modifySlice(a []int) {
	a[0] = 10
	a[1] = 11

	// a = append(a, 12)
	fmt.Printf("before append:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))

	a = append(a, 12, 13, 14)

	fmt.Printf("append1:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))

	a = append(a, 15)

	fmt.Printf("append2:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))

	a = append(a, 16, 17, 18)

	fmt.Printf("append3:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
}

// 切片扩容后，原先切片的底层指向数组的指针会被替换为一个指向新的、长度足够的数组的指针，同时原有的数据也会被拷贝到新的数组地址中。
