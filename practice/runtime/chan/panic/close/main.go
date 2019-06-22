package main

import "sync"

var cnt = 3

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(cnt)

	for i := 1; i <= cnt; i++ {
		j := i
		go func() {
			c <- j
			wg.Done()
		}()
	}

	go func() {//fix: fatal error: all goroutines are asleep - deadlock!
		wg.Wait()
		close(c)
	}()

	for v := range c {
		println(v)
	}
}
