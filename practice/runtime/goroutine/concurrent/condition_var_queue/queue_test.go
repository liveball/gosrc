package main

import (
	"strconv"
	"sync"
	"testing"
)

func TestQueue_Put_Get(t *testing.T) {
	q := NewQueue()

	i := 10
	j := 10
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()

		for i > 0 {
			q.Put(Item{Name: "a" + strconv.Itoa(i), Value: i})
			i--
		}
	}()

	go func() {
		defer wg.Done()

		for j > 0 {
			t.Log(q.Get())
			j--
		}
	}()
	wg.Wait()

	q.Close()
}
