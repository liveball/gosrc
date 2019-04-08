package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	// "runtime"
	"time"
)

func deadloop() {
	for {
	}
}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	go deadloop()
	i := 3
	for {
		time.Sleep(time.Second * 1)
		i--
		fmt.Println("I got scheduled!")
		if i == 0 {
			runtime.GC()
		}
	}
}

//export GOMAXPROCS=1 && go build -o main /data/app/go/src/gosrc/demo/runtime/goroutine/for/main.go && GODEBUG=schedtrace=10000,scheddetail=1 ./main
