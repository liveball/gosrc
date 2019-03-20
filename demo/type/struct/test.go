package main

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
