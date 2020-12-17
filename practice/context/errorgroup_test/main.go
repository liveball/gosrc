package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	//group := new(errgroup.Group)

	group, ctx := errgroup.WithContext(context.Background())

	nums := []int{-1, 0, 1}
	//runtime.GOMAXPROCS(1)
	for _, num := range nums {
		tempNum := num // 子协程中若直接访问num，则可能是同一个变量，所以要用临时变量

		group.Go(func() error {
			if tempNum < 0 {
				return errors.New("tempNum < 0 !!!")
			}

			fmt.Println("tempNum:", tempNum)

			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done")
				return ctx.Err()
			default:
				return nil
			}

		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Get all num successfully!")
	}
}
