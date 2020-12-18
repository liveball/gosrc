package main

import (
	"fmt"
	"net"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(1) //保证单核并发执行，防止ints被多个处理器访问
}

func main() {
	_, _ = net.ResolveTCPAddr("tcp", ":4040")

	// foo := addByShareMemory(10)
	foo := addByShareCommunicate(2)

	fmt.Println(len(foo))
	fmt.Println(foo)

}

func addByShareMemory(n int) []int {
	ints := make([]int, 0, n)

	var mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			mux.Lock()
			ints = append(ints, i)
			mux.Unlock()

			wg.Done()
		}(i)
	}

	wg.Wait()

	return ints
}

func addByShareCommunicate(n int) []int {
	ints := make([]int, 0, n)
	ch := make(chan int, n)

	for i := 0; i < n; i++ {
		go func(c chan int, order int) {
			fmt.Println("input:", order, time.Now())
			c <- order
		}(ch, i)
	}

	for i := range ch {
		fmt.Println("output:", i, time.Now())

		ints = append(ints, i)

		if len(ints) == n {
			break
		}
	}

	close(ch)

	return ints
}
