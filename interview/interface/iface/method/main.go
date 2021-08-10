package main

type Person interface {
	Say() string
}

type Man struct {
}

func (m *Man) Say() string {
	return "Man"
}

func main() {
	var p Person

	m := &Man{}
	p = m
	println(p.Say())
}
