package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	////go handle(ctx, 500*time.Millisecond)
	//go handle(ctx, 1500*time.Millisecond)
	//
	//select {
	//case <-ctx.Done():
	//	fmt.Println("main", ctx.Err())
	//}

	ctx, cancel := context.WithCancel(context.Background())
	go HandleRequest(ctx)

	fmt.Println("it is time to stop all sub goroutine")
	time.Sleep(time.Second * 5)
	
	cancel()

	time.Sleep(time.Second * 5)

}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func HandleRequest(ctx context.Context) {
	go WriteRedis(ctx)
	go WriteDataBase(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandleRequest Done.")
			return

		default:
			fmt.Println("HandleRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDataBase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDataBase Done.")
			return
		default:
			fmt.Println("writeDataBase running")
			time.Sleep(time.Second * 2)
		}
	}
}

func WriteRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(time.Second * 2)
		}
	}
}
