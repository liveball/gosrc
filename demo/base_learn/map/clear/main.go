package main

import (
	"fmt"
	//	"time"
)

type person struct {
	mid   int
	ty    int
	count int
}

func main() {
	batchSize := 2

	per := make(chan *person)

	go func() {
		for i := 1; i <= 3; i++ {
			per <- &person{mid: i, ty: 1, count: i * 10}
		}
		per <- &person{mid: 4, ty: 1, count: 40}

		per <- &person{mid: 5, ty: 1, count: 50}
		per <- &person{mid: 6, ty: 1, count: 60}
		per <- &person{mid: 2, ty: 1, count: 200}
		//per <- &person{mid: 7, ty: 1, count: 700}
		close(per)
	}()

	res := make(map[int]map[int]int)
	for v := range per {
		tgMap := make(map[int]int)
		tgMap[v.ty] = v.count
		res[v.mid] = tgMap
		if len(res) == batchSize {
			fmt.Println(111, res)
			res = nil
			res = make(map[int]map[int]int)
		}
		fmt.Println(222, res)
	}
	//time.Sleep(time.Second)
}
