package main

import "sync"

func main() {
	cnt := 3
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 1; i <= cnt; i++ {
		j := i
		go func() {
			defer wg.Done()
			println(j)
			ch <- j
		}()
	}
	wg.Wait()
	close(ch)

	for v := range ch {
		println(v)
	}
}
