package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go print()
	time.Sleep(time.Second)
}

func print() {
	callers := make([]uintptr, 1024)
	n := runtime.Callers(1, callers)
	for _, pc := range callers[:n] {
		funcPc := runtime.FuncForPC(pc)
		fmt.Println(funcPc.Name())
		fmt.Println(funcPc.FileLine(pc))
	}
}
