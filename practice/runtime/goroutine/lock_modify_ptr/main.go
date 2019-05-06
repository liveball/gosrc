package main

import (
	"sync"
)

var lock sync.Mutex

func a(num *int, wg *sync.WaitGroup) {
	lock.Lock()
	*num++
	// println("a", *num)
	lock.Unlock()

	wg.Done()
}
func main() {
	var num *int
	num = new(int)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go a(num, &wg)
	}
	wg.Wait()
	println("main", *num)
}
