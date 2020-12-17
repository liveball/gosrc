package main

import (
	"runtime/trace"
	"sync"
	"os"
	"log"
)

func main() {
	//runtime.GOMAXPROCS(2)

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			a := 0

			for i := 0; i < 1e6; i++ {
				a += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()
}
