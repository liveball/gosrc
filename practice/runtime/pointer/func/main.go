package main

import "fmt"

func main() {
	f1, f2 := f(), f()
	fmt.Println(*f1, *f2)
}

func f() *int {
	v := 1
	return &v
}
