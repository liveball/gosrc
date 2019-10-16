package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0

	go func() {
		for {
			a++
			//time.Sleep(time.Second*1)
		}
	}()

	time.Sleep(time.Second*1)
	fmt.Println(a)
}

//1、逃逸分析
//2、go程cpu密集，未出让cpu，协程每次去主协程获取寄存器中存储的a内存地址
//3、for循环里加个sleep或者print之类的io操作，应该就不会被优化成取cache了