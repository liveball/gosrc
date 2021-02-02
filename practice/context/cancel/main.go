package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*250))
	defer cancel()

	code, err := apiMock(ctx, time.Millisecond*100)
	if err != nil {
		fmt.Println("mock1:", err)
	}

	code1, err := apiMock(ctx, time.Millisecond*200)
	if err != nil {
		fmt.Println("mock2:", err)
	}

	code2, err := apiMock(ctx, time.Millisecond*10)
	if err != nil {
		fmt.Println("mock3:", err)
	}

	fmt.Println("mock code:", code)
	fmt.Println("mock code1:", code1)
	fmt.Println("mock code2:", code2)

	fmt.Println("耗时：", time.Now().Sub(start))
}

func apiMock(c context.Context, timeout time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		time.Sleep(timeout) //模拟耗时
		ch <- 1
	}()

	var i int
	select {
	case i = <-ch:
		return i, nil
	case <-c.Done():
		return i, c.Err()
	}

}
