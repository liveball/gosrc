package main

import "fmt"

type Vehicle interface {
	Run()
}

type Car struct{}

func (c Car) Run() {
	fmt.Println("car run")
}
