package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	c := make(chan int, 1)

	go func(ch chan int) {
		c <- 1
		close(c)
	}(c)

	select {
	case i, ok := <-c:
		if ok {
			fmt.Println(i)
		}
	case <-time.After(time.Second):
		fmt.Println("c not read")
	}

	fmt.Printf("start: %d Kb\n", m.Alloc/1024)

	runtime.SetFinalizer(&c, func(o interface{}) {
		println("c dead")
	})

	runtime.GC()
	println("read c:", <-c)

	fmt.Printf("end: %d Kb\n", m.Alloc/1024)

	time.Sleep(time.Second * 60)
	// runtime.KeepAlive(c) //确保c指针活着,这样gc完成后，可以观察c指针保持可达状态
}
