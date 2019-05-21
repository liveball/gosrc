package main

import "fmt"

type person struct {
	name string
}

func main() {
	//new
	p := new(person)
	p.name = "aa"
	fmt.Println(p)

	//make  slice map chan
	ps := make([]*person, 0)
	ps = append(ps, p)
	fmt.Println(ps)

	pMap := make(map[string]*person)
	pMap[p.name] = p
	fmt.Println(pMap)

	pChan := make(chan *person)
	go func() {
		pChan <- p
	}()
	fmt.Println(<-pChan)
}
