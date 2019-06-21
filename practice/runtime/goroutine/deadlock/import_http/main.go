package main

// import (
// 	_ "net/http"
// )

func main() {
	println(1)
	<-make(chan int)
}
