package main

import "fmt"

type Ms struct {
	Name string
	Cgs  int
}

func main() {
	hashRun := make(map[string]*Ms)
	mss := []Ms{
		{Name: "m", Cgs: 100},
		{Name: "m", Cgs: 1},
	}

	for _, ms := range mss {
		if h, ok := hashRun[ms.Name]; ok {
			h.Cgs += ms.Cgs
		} else {
			hashRun[ms.Name] = &ms
		}
	}

	for _, ms := range hashRun {
		fmt.Println(ms.Cgs)
	}
}
