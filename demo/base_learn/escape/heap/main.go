package main

type person struct {
	name string
	age  int
}

func foo() *person {
	var p person
	p.name = "pi"
	p.age = 30
	return &p
}

func main() {
	p := foo()
	println(p)
}
