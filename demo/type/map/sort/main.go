package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func main() {
	pm := map[int]*person{
		1: &person{age: 12},
		2: &person{age: 3},
		3: &person{age: 1},
		4: &person{age: 4},
	}

	pms := make([]*person, 0, len(pm))
	for _, v := range pm {
		pms = append(pms, v)
	}
	sort.Slice(pms, func(i, j int) bool {
		return pms[i].age < pms[j].age
	})

	for _, v := range pms {
		fmt.Println(v)
	}
}
