package main

import (
	"runtime"
	"time"
)

func main() {
	println(1 << 20)    //1<<10 1k 10<<10 10k 20<<10 20k 200<<10 200k
	var x [1 << 20]byte //1<<20  1048576  100w不会被回收 10<<20  104857600 1亿被回收

	runtime.SetFinalizer(&x, func(o interface{}) {
		println("x 对象被回收") //大对象也会被立即回收
	})

	time.Sleep(time.Second)
}
