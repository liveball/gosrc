package main

import (
	"fmt"
	"runtime"
	"time"
)

var forcegcperiod int64 = 2 * 60 * 1e9
var ForceGCPeriod = &forcegcperiod

func main() {

	var before, after runtime.MemStats
	//var middle runtime.MemStats
	runtime.ReadMemStats(&before)

	runtime.GC()

	// Make periodic GC run continuously.
	//orig := *ForceGCPeriod
	//*ForceGCPeriod = 0

	fmt.Println("before 堆对象:", before.HeapObjects)

	for i := 1; i < 10000; i++ {
		foo()

		//runtime.GC()

		//runtime.ReadMemStats(&middle)
		//fmt.Println("middle 堆对象:", middle.HeapObjects)
		//time.Sleep(time.Second * 1)
	}

	//runtime.GC()

	//time.Sleep(time.Second * 30)
	time.Sleep(time.Minute * 3)

	runtime.ReadMemStats(&after)
	numsGC := after.NumGC - before.NumGC
	//*ForceGCPeriod = orig

	fmt.Println("after 堆对象:", after.HeapObjects)
	fmt.Println("gc次数:", numsGC)
}

//go:noinline
func foo() {
	ch := make(chan int, 1)
	ch <- 1
}
