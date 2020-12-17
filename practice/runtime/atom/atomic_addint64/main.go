package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var total int64
	sum := 0

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		sum += i

		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&total, int64(i))//不加waitgroup WARNING: DATA RACE
			wg.Done()
		}(i)
	}

	//time.Sleep(1e9)
	wg.Wait()
	fmt.Printf("total:%d, sum:%d\n", total, sum)
}
