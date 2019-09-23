package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
)

func main() {
	f, err := os.OpenFile("heapdump", os.O_RDWR|os.O_CREATE, 0666)
	// f, err := os.Create("heapdump")
	if err != nil {
		panic(err)
	}

	debug.WriteHeapDump(f.Fd())

	data := make([]byte, 10, 10)
	fmt.Println(data)
	runtime.GC()
}
