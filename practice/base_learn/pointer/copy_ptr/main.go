package main

import "fmt"

//Person for person info
type Person struct {
	//Name for name
	Name string
}

func main() {
	var p Person
	p.Name = "小明"

	var i interface{}
	i = p

	if ii, ok := i.(Person); ok {
		ii.Name = "小黑"
		fmt.Println(ii)

		fmt.Println(p == ii)
	}

	fmt.Println(p == i, p.Name, i)
}
