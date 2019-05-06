package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var (
	c = make(chan int)
)

func main() {
	statGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Printf("goroutine num(%d)\n", total)
		}
	}
	go statGroutine()

	go checkTest()

	http.HandleFunc("/", testGo)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func testGo(w http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		c <- 1
		wg.Done()
	}()
	wg.Wait()
	io.WriteString(w, "hello, world!\n")
}

func checkTest() {
	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()
	for {
		select {
		case i := <-c:
			println("i:", i)
		case <-ticker.C:
			time.Sleep(time.Second)
			println("tick...")
		}
	}
}
