package main

import "runtime"

func main() {
	runtime.GOMAXPROCS(10)

	iCh := make(chan int, 1)
	sCh := make(chan string)

	iCh <- 1
	sCh <- "hello"

	select {
	case v := <-iCh:
		println("int chan", v)
	case v := <-sCh:
		println("string chan", v)
	}
}
