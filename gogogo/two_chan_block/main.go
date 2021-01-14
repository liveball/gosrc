package main

import "fmt"

func main() {
	c := make(chan int, 2)
	go bgsweep(c)
	go bgscavenge(c)
	<-c
	<-c
}

func bgsweep(c chan int) {
	c <- 1
	fmt.Println("bgsweep")
}

func bgscavenge(c chan int) {
	c <- 2
	fmt.Println("bgscavenge")
}
