package main

import "fmt"

const b = 123

func main() {
	a := "123"
	//a[1] = 'x'

	ab := []byte(a)
	ab[1] = 'x'

	fmt.Println(a, string(ab))
	//println(&b)
}
