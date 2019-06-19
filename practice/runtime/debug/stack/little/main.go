package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	go print()
	time.Sleep(time.Second)
}

func print() {
	fmt.Println(string(debug.Stack()))
}
