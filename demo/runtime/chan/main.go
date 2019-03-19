package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// lastG()
	// allG()
	// fiveG()
	// fiveWaitG()
	multipleG()
}

// I have multiple goroutines trying to receive on the same channel simultaneously.
// It seems like the last goroutine that starts receiving on the channel gets the value.
// Is this somewhere in the language spec or is it undefined behaviour?
func lastG() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		go func(i int) {
			<-c
			c <- fmt.Sprintf("goroutine %d", i)
		}(i)
	}
	c <- "hi"
	fmt.Println(<-c)
}

// I just realized that it's more complicated than I thought.
// The message gets passed around all the goroutines.
func allG() {
	c := make(chan string)
	for i := 0; i < 5; i++ {
		go func(i int) {
			msg := <-c
			c <- fmt.Sprintf("%s, hi from %d", msg, i)
		}(i)
	}
	c <- "original"
	fmt.Println(<-c)
}

// It creates the five go-routines writing to a single channel, each one writing five times.
// The main go-routine reads all twenty five messages - you may notice that the order they appear in is often not sequential (i.e. the concurrency is evident).

// This example demonstrates a feature of Go channels:
// it is possible to have multiple writers sharing one channel; Go will interleave the messages automatically.

func fiveG() {
	c := make(chan int)
	for i := 1; i <= 5; i++ {
		go func(i int) {
			for v := range c {
				fmt.Printf("count %d from goroutine #%d\n", v, i)
			}
		}(i)
	}
	for i := 1; i <= 25; i++ {
		c <- i
	}
	close(c)
}

func fiveWaitG() {
	c := make(chan int)
	var w sync.WaitGroup
	w.Add(5)

	for i := 1; i <= 5; i++ {
		go func(i int, ci <-chan int) {
			j := 1
			for v := range ci {
				time.Sleep(time.Millisecond)
				fmt.Printf("%d.%d got %d\n", i, j, v)
				j++
			}
			w.Done()
		}(i, c)
	}

	for i := 1; i <= 25; i++ {
		c <- i
	}
	close(c)
	w.Wait()
}

// For multiple goroutine listen on one channel, yes, it's possible.
// the key point is the message itself, you can define some message like that:

type obj struct {
	msg      string
	receiver int
}

func multipleG() {
	ch := make(chan *obj) // both block or non-block are ok
	var wg sync.WaitGroup
	receiver := 25 // specify receiver count

	sender := func() {
		o := &obj{
			msg:      "hello everyone!",
			receiver: receiver,
		}
		ch <- o
	}

	recv := func(idx int) {
		defer wg.Done()
		o := <-ch
		fmt.Printf("%d received at %d\n", idx, o.receiver)
		o.receiver--
		if o.receiver > 0 {
			ch <- o // forward to others
		} else {
			fmt.Printf("last receiver: %d\n", idx)
		}
	}

	go sender()
	for i := 0; i < receiver; i++ {
		wg.Add(1)
		go recv(i)
	}

	wg.Wait()
}
