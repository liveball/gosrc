package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// https://my.oschina.net/u/2336761/blog/2221970

//GOMAXPROCS=1
//$GODEV/bin/go build -o main main.go && GODEBUG=schedtrace=10000,scheddetail=1 ./main

func handle(i interface{}) {
	j, ok := i.(int)
	if !ok {
		return
	}

	fmt.Println(j)

	// fmt.Println(j % 10)
}

//一个p执行的时候，会把当前g放入全局队列
func main() {

	concurrent()
}

func work() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
		}()
		wg.Done()
	}
	wg.Wait()

	// Wait to see the global run queue deplete.
	time.Sleep(3 * time.Second)
}

func test() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("a", i)
			wg.Done()
		}()
	}

	for i := 0; i < 1; i++ {
		go func(j int) {
			fmt.Println("b", j) //runnext 执行4
			wg.Done()
		}(i)
	}
}

func concurrent() {
	// runtime.GOMAXPROCS(1) //Concurrent
	runtime.GOMAXPROCS(2) //Parallel

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		// time.Sleep(1 * time.Microsecond)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
