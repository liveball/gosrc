package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	kf := func() (k chan struct{}) {
		k = make(chan struct{})
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				k <- struct{}{}
				wg.Done()
			}()
		}
		return k
	}

	var g int
	kk := kf()

	go func() {
		wg.Wait()
		close(kk)
	}()

	for {
		_, ok := <-kk
		if !ok {
			return
		}

		g++
		gg := g
		go func() {
			println("gg:", gg)
		}()
	}
}
