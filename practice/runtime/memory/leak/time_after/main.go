package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//leadkMem()

	//testMemStat()

	timer()
}

func leadkMem() {
	ch := make(chan int, 10)

	go func() {
		var i = 1
		for {
			i++
			ch <- i
		}
	}()

	for {
		select {
		case x := <-ch:
			fmt.Println(x)
		case <-time.After(3 * time.Minute):
			fmt.Println(time.Now().Format("2006-01-02 15:16:05"))
		}
	}

}

func testMemStat() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("before, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")
	for i := 0; i < 1000000; i++ {
		time.After(1 * time.Minute)
	}
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("after, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	time.Sleep(10 * time.Second)
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("after 10sec, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")

	time.Sleep(1 * time.Minute)
	runtime.GC()
	runtime.ReadMemStats(&ms)
	fmt.Println("after 3min, have", runtime.NumGoroutine(), "goroutines,", ms.Alloc, "bytes allocated", ms.HeapObjects, "heap object")
}

func timer() {
	ch := make(chan int, 10)

	go func() {
		for {
			ch <- 100
			time.Sleep(time.Second)
		}
	}()

	idleDuration := 1 * time.Second
	idleDelay := time.NewTimer(idleDuration)
	defer idleDelay.Stop()

	for {
		idleDelay.Reset(idleDuration)

		select {
		case x := <-ch:
			fmt.Println(x)
		case c := <-idleDelay.C:
			fmt.Println(c.Format("2006-01-02 15:16:05"))
			//return
		}
	}

}