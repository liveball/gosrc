package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
	"unsafe"
)

var ch = make(chan int)
var ch1 = make(chan int)
var quit chan int // 只开一个信道

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n\r", i)
	}
	ch <- 100 // 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走
}

func say(s string) {
	fmt.Println(s)
	ch <- <-ch1 //ch 等待ch1流出的数据
}

func foo2(id int) {
	quit <- id // ok, finished
}

func test1() {
	loop()
	// go loop()
	time.Sleep(time.Second)
}

func test2() {
	msg := make(chan int, 10)

	statGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Printf("goroutine num(%d)\n", total)
		}
	}
	go statGroutine()

	for i := 0; i < 10; i++ {
		go func(i int) {
			msg <- i
		}(i)
	}

	ticker := time.NewTicker(time.Millisecond * 100)
	q := make([]int, 0)
	f := func() {
		for {
			select {
			case id := <-msg:
				q = append(q, id)
				fmt.Printf("q 扩容：%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&q)))

				fmt.Println(q)
				fmt.Println("------------------")
			case <-ticker.C:
				// go func() {//会引起 data race
				if len(q) == 0 {
					return
				}
				q = q[1:]
				fmt.Printf("q...:(%d)\n", q)
				fmt.Printf("q 缩容：%+v\n", (*reflect.SliceHeader)(unsafe.Pointer(&q)))

				time.Sleep(time.Second * 1) //模拟 api
				// }()
				// default:
				// 	time.Sleep(time.Second * 1)
				// 	fmt.Println("defautl")
			}
		}
	}

	// for i := 0; i < 5; i++ {//模拟并发
	go f()
	// }

	// fmt.Println(<-msg)

	time.Sleep(time.Second * 60)
}

func test3() { //如果不用信道来阻塞主线的话，主线就会过早跑完，loop线都没有机会执行
	go foo()
	go func() {
		ch <- 1 //单线死锁
	}()
	fmt.Printf("直到线程跑完, 取到消息. main在此阻塞住!!! %v\n", <-ch)
}

func test4() { //多线死锁
	go say("Hello")
	<-ch // 堵塞主线
}

func test5() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3//缓冲信道死锁
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

func test6() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	//如果不显式地关闭信道，下面的代码，会报死锁错误的，原因是range不等到信道关闭是不会结束读取的，
	//也就是如果 缓冲信道干涸了，那么range就会阻塞当前goroutine, 所以死锁。
	close(ch)

	for v := range ch {
		fmt.Println(v)
		// if len(ch) <= 0 { // 如果现有数据量为0，跳出循环,防止死锁。
		// 	break
		// }
	}
}

func test7() {
	// 只使用单个无缓冲信道阻塞主线
	count := 10
	quit = make(chan int) // 无缓冲
	for i := 0; i < count; i++ {
		go foo2(i)
	}
	for i := 0; i < count; i++ {
		println(<-quit)
	}
}

func test8() {
	//使用容量为goroutines数量的缓冲信道
	count := 10
	var quit2 chan int
	quit2 = make(chan int, count) // 缓冲
	for i := 0; i < count; i++ {
		quit2 <- i
	}
	close(quit2)
	for v := range quit2 {
		println(v)
		// if len(quit2) <= 0 {
		// 	break
		// }
	}
	println(<-quit2)
	println(<-quit2)
	println(<-quit2)
}

func test9() {
	// Send 1000 integers in 4 goroutines,
	// ensure that we receive what we send.
	chanCap := 100
	c := make(chan int, chanCap)

	const P = 4
	const L = 3
	for p := 0; p < P; p++ {
		go func() { 
			for i := 0; i < L;i++ {
				c <- i
			}
		}()
	}

	done := make(chan map[int]int)
	for p := 0; p < P; p++ {
		go func() {
			recv := make(map[int]int)
			for i := 0; i < L; i++ {
				v := <-c
				fmt.Println(v)
				fmt.Println("-------------")

				recv[v] = recv[v] + 1
			}
			done <- recv
			// fmt.Println(recv)
		}()
	}

	recv := make(map[int]int)
	for p := 0; p < P; p++ {
		for k, v := range <-done {
			recv[k] = recv[k] + v
		}
	}

	// spew.Dump(recv)
	if len(recv) != L {
		fmt.Printf("chan[%d]: received %v values, expected %v", chanCap, len(recv), L)
	}
	for _, v := range recv {
		if v != P {
			fmt.Printf("chan[%d]: received %v values, expected %v", chanCap, v, P)
		}
	}
}
func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	// test5()
	// test6()
	// test7()
	// test8()

	test9()
}
