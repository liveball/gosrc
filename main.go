package main

import "log"

func main() {

	var c chan int
	c = make(chan int)
	go func() {
		c <- 1
	}()

	println(<-c)

	a := make([]int, 0)
	a = append(a, 1)
	log.Println(a)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		// println(x)
		sum += x
		println(sum)
		return sum
	}
}

//func main() { // breakpoint 1
//	ch := make(chan int)
//	go func() {
//		for i := range ch {
//			println(i) // breakpoint 2
//		}
//	}()
//
//	ch <- 1
//
//	wait := make(chan int) // breakpoint 3
//	<-wait
//}
