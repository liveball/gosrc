package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

func main() {
	fmt.Println(debug.SetGCPercent(1), 4<<20)

	// 1
	var dic = make([]byte, 100, 100)
	runtime.SetFinalizer(&dic, func(dic *[]byte) {
		fmt.Println("内存回收1")
	})

	// 立即回收
	runtime.GC()

	// 2
	var s = make([]byte, 100, 100)
	runtime.SetFinalizer(&s, func(dic *[]byte) {
		fmt.Println("内存回收2")
	})

	// 3
	d := make([]byte, 300, 300)
	for i, _ := range d {
		d[i] = 'a' + ([]byte)(strconv.Itoa(i))[0]
	}
	fmt.Println(d)

	time.Sleep(time.Second)
}

// 1处我们创建了一块内存空间100字节，只有我们调用了runtime.GC()立即回收了内存，
// 2处我们又创建了一块100字节的内存，等待回收，
// 当我们执行到3处的时候，创建了一个300字节的内存,已大于垃圾回收剩余内存,所以系统继续立即回收内存。
