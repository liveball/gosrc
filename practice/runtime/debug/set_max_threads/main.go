package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	// 	runtime: program exceeds 1-thread limit
	// fatal error: thread exhaustion

	debug.SetMaxThreads(1)
	go print()

	time.Sleep(time.Second)
}

func print() {
	fmt.Println("1")
}

// 我们把程序的组大可使用的线程（不是协程）数设置为1，如果程序试图超过这个限制,程序就会崩溃，初始设置为10000个线程
// 什么时候会创建新的线程呢?
// 现有的线程阻塞，cgo或者runtime.LockOSThread函数阻塞其他go协程
