package main

import "github.com/fengpf/goroutineid"

func main() {
	id := goroutineid.GetGoID()
	println(id) // 1
}