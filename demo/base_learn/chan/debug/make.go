package main

import (
	"fmt"
	"time"
)

//go tool objdump -s "main\.\w+" chan

//go build -o chan -gcflags="-N -l" chan.go

//go build -o chan -gcflags "-N -l" && GODEBUG=gctrace=1   ./chan

//GODEBUG=gctrace=1 go run  chan.go

func main() {
	// a := make(chan int, 1)
	i := 2
	a := make(chan *int, 1)
	// go func() {
	a <- &i
	// a <- 2
	// }()
	close(a)
	for {
		select {
		case i, ok := <-a:
			if ok {
				fmt.Println("a:", i, "\tStatus:", ok)
			} else {
				fmt.Println("a over")
				return
			}
		default:
			fmt.Println("default over")
			return
		}
	}

	time.Sleep(10 * time.Second)
}

func ChanDebug() {
	var ch2 chan int
	ch2 = make(chan int, 30)

	for i := 0; i < 20; i++ {
		ch2 <- 10000 + i
	}

	//	overwrite ? DeadLock !
	//for i := 0; i < 20; i++ {
	//	ch2 <- 80000 + i
	//}

	for j := 0; j < 15; j++ {
		if i, ok := <-ch2; ok {
			fmt.Println("ch2:", i, "\tStatus:", ok)
		}
	}

	for i := 0; i < 20; i++ {
		ch2 <- 20000 + i
	}

	for {
		select {
		case i, ok := <-ch2:
			fmt.Println("ch2:", i, "\tStatus:", ok)
		default:
			fmt.Println("ch2 over")
			return
		}
	}
}
