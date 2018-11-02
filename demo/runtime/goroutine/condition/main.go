package main

import (
	"fmt"
	"sync"
	"time"
)

type cat struct {
	name string
	mux  sync.Mutex
}

func main() {

	c := &cat{}
	//bug！！！ modify c  by two more goroutine
	//use lock not fix yet
	//but create two cat instances

	var wg sync.WaitGroup

	var fname, sname string
	wg.Add(2)
	go func() {
		c.mux.Lock()
		c.name = "jim"
		c.mux.Unlock()

		fname = getName(c)
		wg.Done()
	}()
	go func() {
		c.mux.Lock()
		c.name = "tom"
		c.mux.Unlock()

		sname = getName(c)
		wg.Done()
	}()

	wg.Wait()

	fmt.Printf("fname: %s sname: %s\n", fname, sname)

}

func getName(c *cat) string {
	time.Sleep(time.Millisecond * 50)
	return c.name
}
