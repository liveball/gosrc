package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"testing"
)

func TestSliceBuffer(t *testing.T) {
	f, err := os.Open("/Users/fpf/Downloads/test.mp3")
	if err != nil {
		t.Fatalf("os.Open error(%v)", err)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		t.Fatalf("f.Stat error(%v)", err)
		return
	}
	if fi.Size() == 0 {
		t.Fatalf("fi.Size() == 0")
		return
	}

	var (
		total   = fi.Size()
		perSize = int(math.Ceil(float64(total) / float64(3)))
		i       = 0
	)

	sum := int64(0)
	for {
		buf := make([]byte, perSize)
		n, e := f.Read(buf[:])
		if e != nil {
			if e == io.EOF {
				break
			}
			err = e
			t.Logf("f.Read error(%v)", err)
			return
		}
		if n != perSize { //最后一片
			buf = buf[0:n]
		}

		i++

		sum += int64(len(buf))

		fmt.Println(n, len(buf))
	}

	fmt.Println(total, total == sum)

}

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
