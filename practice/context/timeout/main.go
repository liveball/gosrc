package main

import (
	"context"
	"gosrc/go/src/fmt"
	"time"
)

func main() {
	start := time.Now()
	getData(context.Background())
	fmt.Println("耗时：", time.Now().Sub(start))
}

func getData(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Duration(time.Millisecond*200))
	defer cancel()

	res, err := api(ctx) // ctx/c
	fmt.Println(res, err)
}

func api(c context.Context) (int, error) {
	ch := make(chan int)
	go func() {
		time.Sleep(time.Millisecond * 500) //模拟耗时
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
