package main

import (
	"fmt"
	"hash/crc32"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count        = 5
	maxProducer  = 5
	maxConsume   = 5
	chanSize     = 1024
	consumeQueue = make([]chan int, maxConsume)

	n = 10000
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)

	// consume(&wg)
	// go produce(&wg)
	// wg.Wait()

	chanProdCons(n, 0, 1000000)
	// time.Sleep(10 * time.Second)
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

func produce(wg *sync.WaitGroup) {
	for i := 0; i < maxProducer; i++ {
		wg.Add(1)
		go func(i int) {
			index := fmt.Sprintf("%02d", i)
			fmt.Printf("queue index:%s\n", index)
			for i := 0; i < count; i++ {
				consumeQueue[shardingQueueIndex(index)] <- i
			}
			wg.Done()
		}(i)
	}
	wg.Done()
}

func consume(wg *sync.WaitGroup) {
	for i := 0; i < maxConsume; i++ {
		ch := make(chan int, chanSize)
		consumeQueue[i] = ch

		wg.Add(1)
		go func(c chan int) {
			for i := range c {
				fmt.Printf("consume at %d\n", i)
				time.Sleep(10 * time.Millisecond)
			}
			wg.Done()
		}(ch)
	}
}

func shardingQueueIndex(name string) (i int) {
	ch := crc32.ChecksumIEEE([]byte(name))
	i = int(ch) % maxConsume
	return
}
