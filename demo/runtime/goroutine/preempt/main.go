package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	count := runtime.GOMAXPROCS(1)
	fmt.Printf("GOMAXPROCS: %d\n", count)

	// goroutine 1
	go func() {
		for {
			// fmt.Println("1")
			// time.Sleep(1 * time.Second)
		}
	}()

	// goroutine 2
	go func() {
		for {
			fmt.Println("2")
			time.Sleep(1 * time.Second)
		}
	}()

	// time.Sleep(20 * time.Second)
	select {}
}

// export GOMAXPROCS=1
// go build -o main -gcflags "-N  -l" /data/app/go/src/gosrc/demo/runtime/goroutine/preempt/main.go && GODEBUG="schedtrace=1000,scheddetail=1" ./main
