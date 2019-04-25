package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func main() {
	ps := make([]*person, 0, 3)
	ps = append(ps, &person{
		age: 3,
	})

	ps = append(ps, &person{
		age: 2,
	})

	ps = append(ps, &person{
		age: 1,
	})

	println("排序前:")
	for _, v := range ps {
		fmt.Println(v)
	}

	sort.Slice(ps, func(i, j int) bool {
		if ps[i] == nil {
			return false
		}
		return ps[i].age < ps[j].age
	})

	println("排序后:")
	for _, v := range ps {
		fmt.Println(v)
	}
}
