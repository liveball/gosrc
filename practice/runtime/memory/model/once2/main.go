package main

import (
	"gosrc/go/src/fmt"
	"sync"
)

var (
	a    string
	done bool
	once sync.Once
)

func setup() {
	a = "hello, world"
	done = true
}

func doprint() {
	if !done {
		once.Do(setup)
	}
	fmt.Println(a)
}

func doprint2() {
	go setup()
	for !done {
	}
	print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	//twoprint()
	doprint2()
}
