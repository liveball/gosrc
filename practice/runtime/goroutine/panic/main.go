package main

import (
	"context"
	"sync"
)

//G for goroutine
type G struct {
	wg  sync.WaitGroup
	ctx context.Context
}

//Go for go func
func (g *G) Go(f func()) {
	g.wg.Add(1)
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
		g.wg.Done()
	}()
	go f()
}

//Wait for go
func (g *G) Wait() {
	g.wg.Wait()
}

func main() {
	//*(*int)(nil) = 0

	g := new(G)
	println(111)
	// ctx, cancel := context.WithCancel(context.Background())
	g.Go(func() {
		panic("goroutine panic")
	})

	g.Go(func() {
		println("!!11")
	})

	g.Wait()
}
