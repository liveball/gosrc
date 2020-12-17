package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))

	return len(p), nil
}

type WordsCounter int

func (wc *WordsCounter) Write(p []byte) (int, error) {
	buf := strings.NewReader(string(p))
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)

	sum := 0
	for s.Scan() {
		fmt.Println(s.Text())
		sum++
	}

	*wc = WordsCounter(sum)
	return sum, nil
}

func main() {
	var c ByteCounter
	n, err := c.Write([]byte("hello"))
	fmt.Println(n, err)

	fmt.Println("c:", c)

	c = 0
	var name = "tom"
	n2, err := fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(n2, err)
	fmt.Println("2 c:", c)

	var wc WordsCounter
	fmt.Fprintf(&wc, "hello world 123")
	fmt.Println(c) //输出 3
}
