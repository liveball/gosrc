package main

import (
	"fmt"
	"sync"
	"unsafe"

	"gosrc/practice/compile/asm/runtime2"
)

func main() {
	//runtime2.Setg(nil)
	fmt.Println("GetGoid:", runtime2.GetGoid())

	atomicTest()
}

func atomicTest() {
	i := 100
	j := 200
	ptr := unsafe.Pointer(&i)
	val := unsafe.Pointer(&j)

	fmt.Println("before StorepNoWB:", ptr, val)
	runtime2.StorepNoWB(ptr, val)
	fmt.Println("after StorepNoWB:", ptr, val)

	a := uint64(300)
	old := uint64(300)
	new := uint64(500)
	//如果*a=old,*a=new,否则啥也不做
	res := runtime2.Cas64(&a, old, new)
	fmt.Println(res, a, old, new)

	k := 0
	b := uint64(0)
	fmt.Println(runtime2.Xadd64(&b, 1))

	var wg sync.WaitGroup
	for k < 500 {
		wg.Add(1)
		go func() {
			fmt.Println(runtime2.Xadd64(&b, 1))

			//b += 1
			//fmt.Println(b)

			wg.Done()
		}()
		k++
	}
	wg.Wait()
	fmt.Println("b:", b)
}
