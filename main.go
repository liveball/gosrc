package main

import (
	"gosrc/go/src/fmt"
	"log"
	"sort"
)

func main() {
	i := 1 << 10
	fmt.Println(i)

	var c chan int
	c = make(chan int)
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)

	a := make([]int, 0)
	a = append(a, 1)
	log.Println(a)

	online := []int{1, 3, 4, 2}
	sort.Slice(online, func(i, j int) bool {
		return online[i] > online[j]
	})
	fmt.Println(online)
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
