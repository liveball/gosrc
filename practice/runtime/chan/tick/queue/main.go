package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	as := New()
	go as.Consume()
	as.run = func(data interface{}) {
		log.Println("consuem:", data)
	}

	i := 0
	for {
		i++
		as.Produce(i)
		time.Sleep(time.Millisecond * 100)
	}
}

// Async struct
type Async struct {
	ch    chan interface{}
	queue []interface{}
	run   func(interface{})
	done  chan struct{}
}

// New new async
func New() *Async {
	return &Async{
		ch:    make(chan interface{}, 1024),
		queue: make([]interface{}, 0, 1024),
		done:  make(chan struct{}),
	}
}

// Consume get data from chan queue
func (a *Async) Consume() {
	t := time.NewTicker(time.Millisecond * 100)
	defer t.Stop()
	for {
		select {
		case data := <-a.ch:
			a.queue = append(a.queue, data)
		case <-t.C:
			if len(a.queue) > 0 {
				data := a.queue[0]
				a.queue = a.queue[1:]
				a.run(data)
			}
		case <-a.done:
			a.queue = a.queue[:0] //清空队列
			log.Println("consuem exit")
			return
		}
	}
}

// Produce put data from chan queue
func (a *Async) Produce(m interface{}) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.ch <- m
	}()
	wg.Wait()
}

// Close queue consume
func (a *Async) Close() {
	a.done <- struct{}{}
}
