package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

var (
	n            = 2000
	wg           sync.WaitGroup
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	now := time.Now()

	//直接处理消息
	directConsume()

	wg.Wait()

	end := time.Now()
	fmt.Printf("%+v\n", end.Sub(now))
}

func directConsume() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		i := 1
		for {
			foo(i)

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
