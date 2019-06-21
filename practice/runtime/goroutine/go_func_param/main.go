package main

import "sync"

var (
	wg  sync.WaitGroup
	cnt = 3
)

func main() {
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		j := i
		go func() {
			println("i=", i, "j=", j)
			wg.Done()
		}()
	}
	wg.Wait()
}
