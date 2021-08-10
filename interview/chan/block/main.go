package main

import "fmt"

//https://juejin.cn/post/6875325172249788429
//go tool compile -N -l -S main.go>main.s

func main(){
	//ch:=make(chan int)

	ch:=make( chan int)
	go func() {
		ch<-1
	}()


	i:=<-ch

	fmt.Println(i)

	//time.Sleep(time.Second*5)
}
