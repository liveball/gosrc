package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	//StopTheWorldDeadlock()

	//testYieldProgress(false)
	//testYieldProgress(true)

	//BlockLocked()
	//
	//time.Sleep(time.Minute)

	fmt.Println(time.Now().AddDate(0, 0, -1).Add(-11 * time.Hour).Format("20060102 15:04:05"))
}

var stop = make(chan bool, 1)

func perpetuumMobile() {
	select {
	case <-stop:
		fmt.Println("stop")
	default:
		//fmt.Println("default")
		go perpetuumMobile()
	}
}

func StopTheWorldDeadlock() {
	maxprocs := runtime.GOMAXPROCS(3)
	log.Println(maxprocs)
	compl := make(chan bool, 2)
	go func() {
		for i := 0; i != 1000; i += 1 {
			runtime.GC()
		}
		compl <- true
	}()
	go func() {
		for i := 0; i != 1000; i += 1 {
			runtime.GOMAXPROCS(3)
		}
		compl <- true
	}()

	go perpetuumMobile()

	<-compl
	<-compl
	stop <- true

	runtime.GOMAXPROCS(maxprocs)
}

func testYieldProgress(locked bool) {
	c := make(chan bool)
	cack := make(chan bool)
	go func() {
		if locked {
			fmt.Println("LockOSThread")
			runtime.LockOSThread()
		}
		for {
			select {
			case <-c:
				cack <- true
				fmt.Println("cack")
				return
			default:
				runtime.Gosched()
			}
		}
	}()
	time.Sleep(10 * time.Millisecond)
	c <- true
	<-cack
}

func BlockLocked() {
	const N = 10
	c := make(chan bool)
	go func() {
		runtime.LockOSThread()
		for i := 0; i < N; i++ {
			c <- true

			time.Sleep(1 * time.Second)
		}
		runtime.UnlockOSThread()
	}()
	for i := 0; i < N; i++ {
		<-c
	}
}