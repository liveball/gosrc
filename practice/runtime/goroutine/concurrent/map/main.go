package main

import (
	"sync"
)

func getData() int {
	return 2
}
func main() {
	var wg sync.WaitGroup
	a := 1
	wg.Add(1)
	go func() {
		// a = getData()
		a = 2
		wg.Done()
	}()
	wg.Wait()

	_ = a
}

// go tool compile -S -race /data/app/go/src/gosrc/practice/runtime/goroutine/concurrent/map/main.go | grep "main.go"
