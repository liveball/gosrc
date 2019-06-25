package main

import (
	_ "net"
)

func main() {
	println(1)
	<-make(chan int)
}
