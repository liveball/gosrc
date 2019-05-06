package main

type person struct {
	name string
	age  int
}

func main() {
	p := new(person)
	p.name = "pi"
	p.age = 30
	println(p)
}
