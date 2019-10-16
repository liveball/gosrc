package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// copy()
		// copyReadFrom()
		copyWriteTo()
		// copyBuffer()
	}
}

func Test_WriteInterface(t *testing.T) {
	var i interface{}
	var buf bytes.Buffer
	i = 1
	buf.WriteString(string(i.(int)))
	fmt.Println(1, buf.String())
	fmt.Println(2, buf.Bytes())
}

// copy()

// goos: darwin
// goarch: amd64
// pkg: readgo/io/buffer
// BenchmarkCopy-4           300000              5205 ns/op           33056 B/op          3 allocs/op
// PASS
// ok      readgo/io/buffer        1.634s

// copyBuffer()

// goos: darwin
// goarch: amd64
// pkg: readgo/io/buffer
// BenchmarkCopy-4          5000000               244 ns/op             288 B/op          2 allocs/op
// PASS
// ok      readgo/io/buffer        1.489s
