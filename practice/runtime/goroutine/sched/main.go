package main

import (
	"fmt"
	"runtime"
)

func g2() {
	sum := 0
	for {
		sum++
		// println(sum)
	}
}

func main() {
	// runtime.GOMAXPROCS(1)
	go g2()

	for {
		runtime.Gosched()
		fmt.Println("main is scheduled!")
	}
}
