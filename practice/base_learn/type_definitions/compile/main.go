package main

import "fmt"

type Celsius float64

var A int

func (c Celsius) String() string  { return fmt.Sprintf("%g°C", c) }
func (c *Celsius) SetF(f float64) { *c = Celsius(f - 32/9*5) }

type foo struct {
}

func main() {
	judge(foo{})
	// judge(&foo{})
}

//assert
func judge(i interface{}) {
	switch i.(type) {
	case struct{}:
		fmt.Println("struct")
	case *struct{}:
		fmt.Println("*struct")
	default:
		fmt.Println("no")
	}
}
