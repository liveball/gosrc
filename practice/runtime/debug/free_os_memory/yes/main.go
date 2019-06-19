package main

import (
	"time"

	"runtime"
	"runtime/debug"
)

func main() {
	var x [1 << 20]byte //10<<20

	runtime.SetFinalizer(&x, func(o interface{}) {
		println("x 对象被回收")
	})
	debug.FreeOSMemory() //小对象立即强制回收

	time.Sleep(time.Second * 1)
}
