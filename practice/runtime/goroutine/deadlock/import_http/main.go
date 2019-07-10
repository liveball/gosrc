package main

import (
	_ "net/http"
)

func main() {
	println(1)
	<-make(chan int)
}

// • 所有 init 函数都在同一个 goroutine 内执行。
// • 所有 init 函数结束后才会执行 main.main 函数。
