package impl

import "fmt"

type Vehicle interface {
	Run()
	Stop()
}

type Car struct{}

func (c Car) Run() {
	fmt.Println("car run")
}
