package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup


func print1(str chan string) {
	fmt.Println("print1")
	time.Sleep(1 * time.Second)
	str <- "print1"
    wg.Done()
}

func print2(str chan string) {
	fmt.Println("print2")
	time.Sleep(1 * time.Second)
	str <- "print2"
	wg.Done()
}

func main() {
	now := time.Now()
	print1Channel := make(chan string)
	print2Channel := make(chan string)

	wg.Add(2)
	go print1(print1Channel)
	print1 := <- print1Channel
	fmt.Println("main", print1)

	go print2(print2Channel)
	print2 := <- print2Channel
	fmt.Println("main",print2)

	wg.Wait()
	since := time.Since(now)
	fmt.Println("耗时--->", since)
}
