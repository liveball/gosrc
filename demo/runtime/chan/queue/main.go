package main

import (
	"fmt"
)

type waitq struct {
	first *sudog
	last  *sudog
}

type sudog struct {

	// isSelect indicates g is participating in a select, so
	// g.selectDone must be CAS'd to win the wake-up race.
	isSelect bool
	next     *sudog
	prev     *sudog

	elem int
}

func main() {

	head := &sudog{
		isSelect: false,
		elem:     0,
	}

	tail := head
	for i := 1; i < 10; i++ {
		tail.next = &sudog{
			isSelect: false,
			elem:     i,
		}
		tail = tail.next
	}

	q := &waitq{
		first: head,
		last:  tail,
	}

	q.enqueue(head)

	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())

}

func (q *waitq) enqueue(sgp *sudog) {
	sgp.next = nil
	x := q.last
	if x == nil {
		sgp.prev = nil
		q.first = sgp
		q.last = sgp
		return
	}
	sgp.prev = x
	x.next = sgp
	q.last = sgp
}

func (q *waitq) dequeue() *sudog {
	for {
		sgp := q.first
		if sgp == nil {
			return nil
		}
		y := sgp.next
		if y == nil {
			q.first = nil
			q.last = nil
		} else {
			y.prev = nil
			q.first = y
			sgp.next = nil // mark as removed (see dequeueSudog)
		}

		// if a goroutine was put on this queue because of a
		// select, there is a small window between the goroutine
		// being woken up by a different case and it grabbing the
		// channel locks. Once it has the lock
		// it removes itself from the queue, so we won't see it after that.
		// We use a flag in the G struct to tell us when someone
		// else has won the race to signal this goroutine but the goroutine
		// hasn't removed itself from the queue yet.
		if sgp.isSelect {
			// if !atomic.Cas(&sgp.g.selectDone, 0, 1) {
			continue
			// }
		}

		return sgp
	}
}
