package main

import "fmt"

type i interface {
	foo()
}

type my struct {
	name string
}

func (m *my) foo() {
	println(m.name)
}

type your struct {
	name string
}

func (y your) foo() {
	println(y.name)

}

func startMyI(name string) i {
	// return my{name: name} //my does not implement i (foo method has pointer receiver)
	return &my{name: name} //ok
}

func startYourI(name string) i {
	// return your{name: name} //ok
	return &your{name: name} //ok
}

func main() {
	fmt.Println(startMyI("startMyI...."))
	fmt.Println(startYourI("startYourI...."))
}
