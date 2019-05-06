package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// go test -bench . -benchmem -gcflags "-N -l"

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}

// goos: darwin
// goarch: amd64
// pkg: readgo/baderror/slice
// BenchmarkArray-4         1000000              2425 ns/op               0 B/op          0 allocs/op
// BenchmarkSlice-4          300000              4267 ns/op            8192 B/op          1 allocs/op
// PASS
// ok      readgo/baderror/slice   3.786s

// 解释一下上述结果，在测试 Array 的时候，用的是4核，循环次数是1000000，平均每次执行时间是2425 ns，每次执行堆上分配内存总量是0，分配次数也是0 。

// 而切片的结果就“差”一点，同样也是用的是4核，循环次数是300000，平均每次执行时间是4267 ns，但是每次执行一次，堆上分配内存总量是8192，分配次数也是1 。

// 这样对比看来，并非所有时候都适合用切片代替数组，因为切片底层数组可能会在堆上分配内存，而且小数组在栈上拷贝的消耗也未必比 make 消耗大。

func nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			println(i, v)
			s[i] = v
			i++
		}
	}
	fmt.Println(i, s)
	return s[:i]
}

func TestNonempty(t *testing.T) {
	a := []string{"one", "", "three"}
	b := nonempty(a)
	fmt.Println(a) //a 被改变le
	fmt.Println(b)
	fmt.Printf("%#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&a)),
		(*reflect.SliceHeader)(unsafe.Pointer(&b)),
	)
}

func TestSliceIndex(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := a[2:3]              //最后一个索引不指定默认是数组长度
	fmt.Println(b[0], &a[2]) //b的长度是1，容量是3，所以b只能取b[0]

	c := a[1:3:4]                  //指定最后一个索引为数组长度4
	fmt.Println(c[0], c[1], &a[1]) //c的长度是2，容量是3， 所以c只能取c[0],c[1]
	fmt.Printf("%#v, %#v, %#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&a)),
		(*reflect.SliceHeader)(unsafe.Pointer(&b)),
		(*reflect.SliceHeader)(unsafe.Pointer(&c)),
	)
}
