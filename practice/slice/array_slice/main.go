package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func testSlice() {
	var array [10]int
	var slice = array[5:6]

	fmt.Println("len of slice:", len(slice))
	fmt.Println("cap of slice:", cap(slice))
	fmt.Println(&slice[0] == &array[5])
}

func main() {
	//testSlice()

	testAddElement()

	//testOrderLen()

	//test()

	sliceA := make([]int, 5, 10)
	sliceB := sliceA[0:5]
	fmt.Printf("sliceA addr(%#v) len(%v) cap(%v)\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&sliceA)), len(sliceA), cap(sliceA))
	fmt.Printf("sliceB addr(%#v) len(%v) cap(%v)\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&sliceB)), len(sliceB), cap(sliceB))
}

func testOrderLen() {
	orderLen := 5
	order := make([]int, 2*orderLen)

	for i := 0; i < 10; i++ {
		order[i] = i
	}

	pollorder := order[:orderLen:orderLen]            //取order 的前半部分
	lockorder := order[orderLen:][:orderLen:orderLen] //取order 的后半部分

	fmt.Println("len(pollorder) cap(pollorder)", len(pollorder), cap(pollorder), pollorder)
	fmt.Println("len(lockorder) cap(lockorder)", len(lockorder), cap(lockorder), lockorder)
}

func AddElement(slice []int, e int) []int {
	fmt.Printf("after call slice addr(%#v) len(%v) cap(%v)\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&slice)), len(slice), cap(slice))

	return append(slice, e)
}

func testAddElement() {
	var slice []int

	slice = append(slice, 1, 2, 3)
	fmt.Println("slice:", len(slice), cap(slice))

	fmt.Printf("before call slice addr(%#v) len(%v) cap(%v)\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&slice)), len(slice), cap(slice))

	newSlice := AddElement(slice, 4)
	fmt.Println("newSlice:", len(newSlice), cap(newSlice))
	fmt.Println(&slice[0] == &newSlice[0])

	newSlice1 := append(newSlice, 5)
	fmt.Println("newSlice1:", len(newSlice1), cap(newSlice1))
	fmt.Println(&newSlice[0] == &newSlice1[0])

	newSlice2 := append(newSlice1, 6)
	fmt.Println("newSlice1:", len(newSlice2), cap(newSlice2))
	fmt.Println(&newSlice1[0] == &newSlice2[0])
}

func test() {
	var arr [8]int = [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("arr %#v\n",
		unsafe.Pointer(&arr),
	)

	fmt.Printf("arr[0] %#v\n",
		unsafe.Pointer(&arr[0]),
	)

	fmt.Printf("arr[1] %#v\n",
		unsafe.Pointer(&arr[1]),
	)

	var s1 []int
	s1 = arr[1:4]
	c1 := cap(s1)
	fmt.Println("s1_len", len(s1))
	fmt.Println("s1_cap", c1)

	fmt.Printf("s1 %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&s1)),
	)

	var s2 = arr[2:4]
	fmt.Println("s2_len", len(s2))
	fmt.Println("s2_cap", cap(s2))

	fmt.Printf("s2 %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&s2)),
	)
}
