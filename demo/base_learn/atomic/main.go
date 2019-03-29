package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var a uint64

	cnt := 5
	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			atomic.AddUint64(&a, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(a)
}
