package main

import (
	"fmt"
	"hash/crc32"
	"os"
	"runtime"
	"runtime/trace"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	n            = 2000
	maxConsume   = 10
	chanSize     = 1024
	consumeQueue = make([]chan int, maxConsume)
	wg           sync.WaitGroup
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	now := time.Now()

	//异步并发处理
	doFoo()
	asyncConsume()

	wg.Wait()

	end := time.Now()
	fmt.Printf("%+v\n", end.Sub(now))

	//chanProdCons(n, 0, 1000000)
	// time.Sleep(10 * time.Second)
}


func doFoo() {
	for i := 0; i < maxConsume; i++ {
		ch := make(chan int, chanSize)
		consumeQueue[i] = ch

		c := ch
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range c {
				foo(i)
			}
		}()
	}
}


func asyncConsume() {
	wg.Add(1)
	go func() {
		defer func() {
			for _, c := range consumeQueue {
				close(c)
			}
			wg.Done()
		}()

		i := 0
		for {
			index := shardingQueueIndexByChecksum(strconv.Itoa(i), maxConsume)
			//index := shardingQueueIndex(i, maxConsume)
			//println("current i:", i, "chan index:", index)

			consumeQueue[index] <- i

			i++

			if i == n {
				fmt.Printf("last msg at %d\n", i)
				break
			}
		}
	}()
}

func foo(i int) {
	//fmt.Printf("consume at %d\n", i)
	time.Sleep(2 * time.Millisecond)
}

func shardingQueueIndex(i, ql int) int {
	return i % ql
}

func shardingQueueIndexByChecksum(name string, ql int) (i int) {
	ch := crc32.ChecksumIEEE([]byte(name))
	i = int(ch) % ql
	return
}

func chanProdCons(n int, chanSize, localWork int) {
	const CallsPerSched = 1000
	procs := runtime.GOMAXPROCS(-1)
	N := int32(n / CallsPerSched)
	c := make(chan bool, 1*procs)
	myc := make(chan int, chanSize)
	for p := 0; p < procs; p++ {
		go func() {
			foo := 0
			for atomic.AddInt32(&N, -1) >= 0 {
				for g := 0; g < CallsPerSched; g++ {
					for i := 0; i < localWork; i++ {
						// foo *= 2
						// foo /= 2

						if i == localWork {
							foo = localWork
						}
						println(i)
					}
					myc <- 1
				}
			}
			myc <- 0
			c <- foo == 1000000
		}()
		// go func() {
		// 	for {
		// 		v := <-myc
		// 		if v == 0 {
		// 			break
		// 		}
		// 		for i := 0; i < localWork; i++ {
		// 			println(i)
		// 		}
		// 	}
		// 	// c <- foo == 42
		// }()
	}
	for p := 0; p < procs; p++ {
		<-c
		// <-c
	}
}
