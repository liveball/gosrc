package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())

	ctx, _ := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		j := i

		go func() {
			for {
				select {
				case <-ctx.Done():
					break
				}

				time.Sleep(time.Millisecond * 100)
			}

			fmt.Println("context cancelled ", j)
		}()
	}

	//cancel()

	time.Sleep(time.Second * 1000)

	fmt.Println("main exit ")
}
