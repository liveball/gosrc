package main

import (
	"fmt"
	"sync"
	"time"

	"go1.11.1/demo/type/chan/batcher/tsdb"
)

func main() {
	batchSize := 0
	batcher := tsdb.NewPointBatcher(batchSize, 0, time.Hour)
	if batcher == nil {
		fmt.Printf("failed to create batcher for size test")
	}

	batcher.Start()

	var wg sync.WaitGroup
	wg.Add(2)
	var p int
	go func() {
		for i := 0; i < batchSize; i++ {
			println(i)
			p = i
			batcher.In() <- p
		}
		wg.Done()
	}()

	go func() {
		batch := <-batcher.Out()
		fmt.Println(batch)
		if len(batch) != batchSize {
			fmt.Printf("received batch has incorrect length exp %d, got %d", batchSize, len(batch))
		}
		wg.Done()
	}()
	wg.Wait()
}
