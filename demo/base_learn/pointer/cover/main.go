package main

import (
	"github.com/davecgh/go-spew/spew"
)

type person struct {
	name string
}

func main() {
	var (
		p = new(person)
	)
	spew.Dump(p)
	p = getPerson()
	spew.Dump(p)
}

func getPerson() *person {
	return &person{name: "aaa"}
}
