package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

// https://www.kancloud.cn/mutouzhang/go/598654
func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()         //1
		defer v1.mu.Unlock() //2

		time.Sleep(2 * time.Second) //3
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
