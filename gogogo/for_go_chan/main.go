package main

import (
	"fmt"
	"runtime"
	"time"
)

var (
	ch = make(chan int)
)

func main() {

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	go func() {
		for {

			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			fmt.Println("goroutine num:", runtime.NumGoroutine(), m.Alloc/1024, m.HeapAlloc/1024, m.HeapObjects/1024)
			time.Sleep(time.Microsecond * 500)
		}
	}()

	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Println("ch close")
				return
			}

			go func(i int) {
				fmt.Println("i--------->", i)
				// time.Sleep(time.Microsecond * 1000)
			}(v)
		}
	}()

	var i int
	for {
		i++
		ch <- i

		// if i == 50 {
		// 	close(ch)
		// }
		time.Sleep(time.Microsecond * 500)
	}
}
