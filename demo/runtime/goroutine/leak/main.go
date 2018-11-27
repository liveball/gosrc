package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
	"time"

	"context"
)

func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {
	stack := debug.Stack()
	w.Write(stack)
	pprof.Lookup("goroutine").WriteTo(w, 2)
}

func main() {
	http.HandleFunc("/stack", getStackTraceHandler)

	// leak()
	// noleak()

	// leak2()
	// noleak2()

	// masterWorkLeak()

	httpLeak()

	time.Sleep(time.Second * 60)
}

func leak() {
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}

func noleak() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := func(ctx context.Context) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					return
				case ch <- i:
				}
			}
		}()
		return ch
	}(ctx)

	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			cancel()
			break
		}
	}
}

func leak2() {
	newRandStream := func() <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)
			// 死循环：不断向channel中放数据，直到阻塞
			for {
				randStream <- rand.Int()
			}
		}()

		return randStream
	}

	randStream := newRandStream()
	fmt.Println("3 random ints:")

	// 只消耗3个数据，然后去做其他的事情，此时生产者阻塞，
	// 若主goroutine不处理生产者goroutine，则就产生了泄露
	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	fmt.Fprintf(os.Stderr, "NumGoroutine:%d\n", runtime.NumGoroutine())
	time.Sleep(10e9)
	fmt.Fprintf(os.Stderr, "NumGoroutine:%d\n", runtime.NumGoroutine())
}

func noleak2() {
	newRandStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)

		go func() {
			defer fmt.Println("newRandStream closure exited.")
			defer close(randStream)

			for {
				select {
				case randStream <- rand.Int():
				case <-done: // 得到通知，结束自己
					return
				}
			}
		}()

		return randStream
	}

	done := make(chan interface{})
	randStream := newRandStream(done)
	fmt.Println("3 random ints:")

	for i := 1; i <= 3; i++ {
		fmt.Printf("%d: %d\n", i, <-randStream)
	}

	// 通知子协程结束自己
	done <- struct{}{}
	// close(done)
	// Simulate ongoing work
	time.Sleep(1 * time.Second)
}

// function to add an array of numbers.
func workerAdder(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// writes the sum to the go routines.
	c <- sum // send sum to c
	fmt.Println("end")
}

func masterWorkLeak() {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	// spin up a goroutine.
	go workerAdder(s[:len(s)/2], c1)
	// spin up a goroutine.
	go workerAdder(s[len(s)/2:], c2)

	//不管在任何情况下，都必须要有协程能够读写channel，让协程不会阻塞
	//x, y := <-c1, <-c2 // receive from c1 aND C2

	x, _ := <-c1
	// 输出从channel获取到的值
	fmt.Println(x)

	fmt.Fprintf(os.Stderr, "NumGoroutine:%d\n", runtime.NumGoroutine())
	time.Sleep(10e9)
	fmt.Fprintf(os.Stderr, "NumGoroutine:%d\n", runtime.NumGoroutine())
}

// 把数组s中的数字加起来
func sumInt(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// HTTP handler for /sum
func sumConcurrent2(w http.ResponseWriter, r *http.Request) {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)

	go sumInt(s[:len(s)/2], c1)
	go sumInt(s[len(s)/2:], c2)

	// 这里故意不在c2中读取数据，导致向c2写数据的协程阻塞。
	// x := <-c1
	// fmt.Fprintf(w, strconv.Itoa(x)+"\n")

	//不管在任何情况下，都必须要有协程能够读写channel，让协程不会阻塞
	x, y := <-c1, <-c2
	// write the response.
	fmt.Fprintf(w, strconv.Itoa(x+y)+"\n")

}

func httpLeak() {
	StasticGroutine := func() {
		for {
			time.Sleep(1e9)
			total := runtime.NumGoroutine()
			fmt.Println("NumGoroutine:", total)
		}
	}

	go StasticGroutine()

	http.HandleFunc("/sum", sumConcurrent2)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
