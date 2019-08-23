package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	_ "net/http/pprof"
)

var wg sync.WaitGroup

func print1Func(str chan string) {
	fmt.Println("print1")
	time.Sleep(1 * time.Second)
	str <- "print1"
	wg.Done()
}

func print2Func(str chan string) {
	fmt.Println("print2")
	time.Sleep(1 * time.Second)
	str <- "print2"
	wg.Done()
}

//func print1Func() (str string) {
//	str = "print1"
//	fmt.Println(str)
//	time.Sleep(1 * time.Second)
//	wg.Done()
//	return
//}
//
//func print2Func() (str string) {
//	str = "print2"
//	fmt.Println(str)
//	time.Sleep(1 * time.Second)
//	wg.Done()
//	return
//}

func main() {
	now := time.Now()

	var print1, print2 string
	print1Channel := make(chan string)
	print2Channel := make(chan string)

	wg.Add(4)
	go print1Func(print1Channel)
	go print2Func(print2Channel)

	go func() {
		print1 = <-print1Channel
		wg.Done()
	}()
	go func() {
		print1 = <-print2Channel
		wg.Done()
	}()

	//go func() {
	//	print1=print1Func()
	//}()
	//go func() {
	//	print2= print2Func()
	//}()
	wg.Wait()

	fmt.Println("main", print1)
	fmt.Println("main", print2)

	since := time.Since(now)
	fmt.Println("耗时--->", since)

	http.ListenAndServe(":9000", nil)
}
