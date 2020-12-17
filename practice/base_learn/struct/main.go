package main

import (
	"fmt"

	"gosrc/go/src/encoding/json"
)

type my struct {
	I int `json:"i" default:"-1"`
}

func (m *my) change() {
	c := &my{3}
	m = c
}

func (m my) change2() {
	// c := &my{4} //cannot use c (type *my) as type my in assignment
	c := my{4}
	m = c
}

func main() {
	m0 := my{}
	fmt.Println(m0)
	a, err := json.Marshal(m0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(a))

	m1 := my{1}
	m1.change()
	fmt.Println(m1.I) //1
	m1.change2()
	fmt.Println(m1.I) //1

	m2 := &my{2}
	m2.change()
	fmt.Println(m2.I) //2
	m2.change2()
	fmt.Println(m2.I) //2
}

//Person for one people
type Person struct {
	Name string
	Age  int
}

func mynew() {
	rect1 := new(Person)
	rect1.Name = "xxx"
	rect1.Age = 22
	fmt.Printf("%v  %T  %v \n", rect1, rect1, *rect1)

	rect2 := &Person{"阿呆", 25}
	fmt.Printf("%v  %T  %v \n", rect2, rect2, *rect2)

	rect3 := Person{"小明", 26}
	fmt.Printf("%v  %T\n", rect3, rect3)
}

func modify(p Person) {
	p.Name = "tom"
}
