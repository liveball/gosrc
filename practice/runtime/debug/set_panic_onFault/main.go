package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	go print()
	time.Sleep(time.Second)

	fmt.Println("ddd")
}

func print() {
	defer func() { recover() }()
	fmt.Println(debug.SetPanicOnFault(true))
	var s *int = nil
	*s = 34
}
