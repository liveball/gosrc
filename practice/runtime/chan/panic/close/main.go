package main

import "sync"

func main() {
	cnt := 3
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		j := i
		go func() {
			defer wg.Done()
			ch <- j
		}()
	}
	close(ch)
	wg.Wait()

	for v := range ch {
		println(v)
	}
}
