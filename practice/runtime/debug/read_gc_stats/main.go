package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

// type GCStats struct {
//     LastGC         time.Time       // 最近一次垃圾收集的时间
//     NumGC          int64           // 垃圾收集的次数
//     PauseTotal     time.Duration   // 所有暂停收集垃圾消耗的总时间
//     Pause          []time.Duration // 每次暂停收集垃圾的消耗的时间
//     PauseQuantiles []time.Duration
// }

func main() {
	data := make([]byte, 1000, 1000)
	println(data)
	runtime.GC()

	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	fmt.Println(stats.NumGC)
	fmt.Println(stats.LastGC)
	fmt.Println(stats.Pause)
	fmt.Println(stats.PauseTotal)
	fmt.Println(stats.PauseEnd)
}
