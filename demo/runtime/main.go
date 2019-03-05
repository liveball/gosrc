package main

import "sync"

// go tool compile -N -l -S main.go > main.s

//GODEBUG=schedtrace=10000,scheddetail=1 ./main
func main() {
	println(1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		println(2)
		wg.Done()
	}()

	wg.Wait()
}
