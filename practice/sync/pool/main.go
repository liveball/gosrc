package main

import "sync"

var bufferPool = sync.Pool{
	New: func() interface{} {
		var b []byte
		return b
	},
}

func newBuffer(size int) []byte {
	b := bufferPool.Get().([]byte)
	if cap(b) < size {
		doublecap := 2 * cap(b)
		if doublecap > size {
			return make([]byte, size, doublecap)
		}
		return make([]byte, size)
	}
	return b[:size]
}

func freeBuffer(b []byte) {
	bufferPool.Put(b[:0])
}



