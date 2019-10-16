package main

import "fmt"

type T struct {
	msg string
}

var g *T

func setup() {
	t := new(T)
	t.msg = "hello, world"
	g = t
}

func main() {
	go setup()
	for g == nil {
		fmt.Println("g is nil")
	}
	fmt.Println(g.msg)
}
