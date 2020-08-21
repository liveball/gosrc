package main

import (
	"sync"
)

type Item struct {
	Name  string
	Value int
}

type Queue struct {
	mu sync.Mutex

	itmes []Item
	cond  sync.Cond

	closed bool
}

func NewQueue() *Queue {
	q := new(Queue)
	q.cond.L = &q.mu
	return q
}

func (q *Queue) Get() Item {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.itmes) == 0 {
		q.cond.Wait()
	}

	item := q.itmes[0]
	q.itmes = q.itmes[1:]
	return item
}

func (q *Queue) Put(item Item) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.itmes = append(q.itmes, item)
	q.cond.Signal()
}

func (q *Queue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.closed = true
	q.cond.Broadcast()
}
