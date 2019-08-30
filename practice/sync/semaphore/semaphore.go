package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const maxWorkers = 5

func main() {
	go func() {
		for {
			log.Println("goroutine:", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	sem := semaphore.NewWeighted(maxWorkers)
	g, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < 50; i++ {
		i := i
		//log.Printf("executing %d\n", i)
		err := sem.Acquire(ctx, 1)
		if err != nil {
			return
		}

		g.Go(func() error {
			defer sem.Release(1)

			// do work
			time.Sleep(1 * time.Second)
			log.Printf("finished %+v\n", i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Printf("g.Wait() err = %+v\n", err)
	}

	log.Println("done!")
}
