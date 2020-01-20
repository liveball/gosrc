package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

var (
	cnt = 20
)

func main() {
	stasticGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Println("NumGoroutine:", total)
		}
	}
	go stasticGroutine()

	for cnt > 0 {
		handle(context.Background(), cnt)
		cnt--
	}

}

func handle(ctx context.Context, cnt int) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ch := make(chan struct{}, 0)
	go func() {
		run(ctx, "run")

		ch <- struct{}{}

		select {
		case <-ctx.Done():
			fmt.Printf("goroutine(%d) exit\n", cnt)
			return
			//default:
			//	fmt.Println("default")
		}
	}()


	select {
	case <-ch:
		fmt.Println("ch")
		close(ch)
		//case <-ctx.Done():
		//fmt.Println("ctx.Done() 2222")
		//default:
		//	fmt.Println("default")
	}

	fmt.Println("handle at", cnt)
}

func run(ctx context.Context, name string) {
	fmt.Println("todo", name)
	time.Sleep(5 * time.Second)
}
