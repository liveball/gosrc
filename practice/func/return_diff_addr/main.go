package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 100
	return func() int {
		// fmt.Printf("%v\n", i)
		i += 1
		return i
	}
}

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		// println(x)
		sum += x
		println(sum)
		return sum
	}
}

func main() {
	// a := []int{1, 2, 3}
	// for _, v := range a {
	// 	println(v)
	// 	// defer p(v)
	// 	defer func() {
	// 		fmt.Printf("%p\n", &v)
	// 		println(v)
	// 	}()
	// }

	println("\ntest intSeq--------\n")

	//We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()
	//See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt,intSeq(),intSeq())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	println("\ntest addr--------\n")

	pos:= addr()
	// _= addr()

	fmt.Println(pos)
	for i := 1; i <=3; i++ {
		pos(i)
	}
}
