package main

import "fmt"

type my struct {
	I int
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
	m1 := my{1}
	m1.change()
	println(m1.I) //1
	m1.change2()
	println(m1.I) //1

	m2 := &my{2}
	m2.change()
	println(m2.I) //2
	m2.change2()
	println(m2.I) //2
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
