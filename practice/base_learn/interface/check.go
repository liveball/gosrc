package main

import "fmt"

type person interface {
	Speak()
	Eat()
}

type student struct {
	Name string
}

func (s *student) Speak() {
	fmt.Println(s.Name + " is a student")
}

func (s *student) Eat() {
	fmt.Println(s.Name + " like eat apple")
}

var _ person = &student{}

func main() {
	s1 := &student{"xiaowang"}
	s1.Speak()

	var s2 person = &student{"xiaohan"}
	s2.Eat()
}
