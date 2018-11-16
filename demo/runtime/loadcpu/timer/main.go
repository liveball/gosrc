package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// testTimer()
	testSelect()
}

func testTimer() {
	// cpu 100%
	//场景1：
	// for {
	// 	select {
	// 	case <-time.After(1 * time.Second):
	// 		fmt.Println("hello timer")
	// 	}
	// }

	//场景2：
	for {
		select {
		case <-time.Tick(10 * time.Microsecond):
			fmt.Println("hello, tick")
		}
	}

	//正确姿势
	// tick := time.Tick(10 * time.Microsecond)
	// for {
	// 	select {
	// 	case <-tick:
	// 		fmt.Printf("hello, tick 2")
	// 	}
	// }

}

func testSelect() {
	quit := make(chan bool)
	for i := 0; i != runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-quit:
					break
				default:
					//如果quit没有数据，则默认走default造成死循环
					time.Sleep(time.Second)
				}
			}
		}()
	}

	time.Sleep(time.Second * 15)
	for i := 0; i != runtime.NumCPU(); i++ {
		quit <- true
	}
}
