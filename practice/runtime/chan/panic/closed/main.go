package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan struct{})
	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(done)

		println(done)
		_, ok := <-done
		if !ok {
			println("close")
		}
	}()


	wg.Add(100)
	cnt:=0
	for i := 0; i < 100; i++ {
		j := i
		go func() {
			println(fmt.Sprintf("%02d", j))
			cnt++
			wg.Done()
		}()
	}

   <-done
   println(cnt)
}
