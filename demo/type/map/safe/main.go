package main

import (
	"fmt"
	"sync"
)

var (
	myMap = make(map[int]int)
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	tmp := make(map[int]int)
	go func() {
		// myMap[1] = 1
		tmp[1] = 1
		wg.Done()
	}()

	myMap = tmp

	go func() {
		// for _, v := range myMap {
		// 	println(v, 1111)
		// }

		fmt.Printf("myMap=%+v\n", myMap)
		wg.Done()
	}()

	wg.Wait()
}
