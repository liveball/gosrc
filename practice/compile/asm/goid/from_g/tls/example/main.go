package main

import (
	"fmt"
	"sync"
	"unsafe"

	"gosrc/practice/compile/asm/goid/from_g/tls"
)

func main() {
	//tls.Setg(nil)
	fmt.Println("GetGoid:", tls.GetGoid())

	str := "addsdsdsds"
	fmt.Println("str len:", tls.StringLen(str))
	fmt.Println("str out:", tls.String(str))

	fmt.Println(1 & 1)

	stu := tls.Student{Name: "小韩", Age: 28}
	fmt.Println("stu:", tls.NewStudent(stu))

	nStu := tls.NewStudentPtr(&stu)
	fmt.Println("nStu1:", *nStu)
	nStu.Name = "小强"
	nStu.Age = 30
	fmt.Println("nStu2:", *nStu)

	tls.UpStudentPtr(nStu, "小王", 40)
	fmt.Println("nStu up:", nStu)

	fmt.Println(tls.StudentName(nStu))
	atomicTest()
}

func atomicTest() {
	i := 100
	j := 200
	ptr := unsafe.Pointer(&i)
	val := unsafe.Pointer(&j)

	fmt.Println("before StorepNoWB:", ptr, val)
	tls.StorepNoWB(ptr, val)
	fmt.Println("after StorepNoWB:", ptr, val)

	a := uint64(300)
	old := uint64(300)
	new := uint64(500)
	//如果*a=old,*a=new,否则啥也不做
	res := tls.Cas64(&a, old, new)
	fmt.Println(res, a, old, new)

	k := 0
	b := uint64(0)
	fmt.Println(tls.Xadd64(&b, 1))

	var wg sync.WaitGroup
	for k < 500 {
		wg.Add(1)
		go func() {
			fmt.Println(tls.Xadd64(&b, 1))

			//b += 1
			//fmt.Println(b)

			wg.Done()
		}()
		k++
	}
	wg.Wait()
	fmt.Println("b:", b)
}
