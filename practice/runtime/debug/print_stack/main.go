package main

import (
	"time"

	"runtime/debug"
)

func main() {
	go print()
	time.Sleep(time.Second)
}

func print() {
	debug.PrintStack()
}
