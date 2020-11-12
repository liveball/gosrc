package main

import (
	"os"
	"log"
	"runtime/trace"
)

func main() {
	if err := trace.Start(os.Stderr); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "aaaa"
	}()

	<-ch
}
