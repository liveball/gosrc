package main

// #include <unistd.h>
import "C"

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			C.sleep(1)                 // 测试1 使用cgo sleep的情况下，可以看到大量的闲置M
			//time.Sleep(time.Second)  // 测试2 类似netpoll的优化，仅阻塞G，不阻塞M，不会创建大量线程
			wg.Done()
		}()
	}
	wg.Wait()
	println("done!")
	time.Sleep(time.Second * 3)
}

//GODEBUG="schedtrace=1000" ./main
//SCHED 0ms: gomaxprocs=12 idleprocs=11 threads=5 spinningthreads=0 idlethreads=2 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//SCHED 1008ms: gomaxprocs=12 idleprocs=12 threads=1004 spinningthreads=0 idlethreads=186 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//done!
//SCHED 2015ms: gomaxprocs=12 idleprocs=12 threads=1004 spinningthreads=0 idlethreads=1000 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//SCHED 3017ms: gomaxprocs=12 idleprocs=12 threads=1004 spinningthreads=0 idlethreads=1000 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//SCHED 4026ms: gomaxprocs=12 idleprocs=12 threads=1004 spinningthreads=0 idlethreads=1000 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]

//GODEBUG="schedtrace=1000" ./main
//SCHED 0ms: gomaxprocs=12 idleprocs=9 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//SCHED 1002ms: gomaxprocs=12 idleprocs=12 threads=11 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//done!
//SCHED 2013ms: gomaxprocs=12 idleprocs=12 threads=12 spinningthreads=0 idlethreads=10 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
//SCHED 3017ms: gomaxprocs=12 idleprocs=12 threads=12 spinningthreads=0 idlethreads=10 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0]
