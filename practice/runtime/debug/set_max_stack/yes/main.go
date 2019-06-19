package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {

	// runtime: goroutine stack exceeds 1-byte limit
	// fatal error: stack overflow

	// 默认的设置32位系统是250MB,64位为1GB
	fmt.Println(debug.SetMaxStack(1)) //查看到默认系统为1000 000 000 字节

	for i := 0; i < 10; i++ {
		go print(i)
	}

	time.Sleep(time.Second)
}

func print(i int) {
	fmt.Println(i)
}
