package main

import (
"fmt"
"sync"
)

func main() {
	var wg sync.WaitGroup

	//cnt := 3

	//for i := 0; i < cnt; i++ {
	//	wg.Add(1)
	//	//fmt.Println("i:", i, &i)
	//
	//	j:=i
	//	go func() {
	//		fmt.Println("j:", j, &j)
	//		wg.Done()
	//	}()
	//}

	//a := []int{1, 2, 3}
	//
	//for _, v := range a {
	//	wg.Add(1)
	//	//fmt.Println("i:", i, &i)
	//
	//	go func() {
	//		fmt.Println("v:", v, &v)
	//		wg.Done()
	//	}()
	//}
	//
	//wg.Wait()


	a := map[int]int{1:1, 2:2, 3:3}

	for _, v := range a {
		wg.Add(1)
		//fmt.Println("i:", i, &i)

		go func() {
			fmt.Println("v:", v, &v)
			wg.Done()
		}()
	}

	wg.Wait()

	b:=[]int{}

	fmt.Println(len(b))
}

