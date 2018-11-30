package main

import "fmt"

func main() {
	// x := struct {
	// }{}

	// x := new(struct { // x size ?
	// 	_ struct{}
	// })

	x := &struct {
	}{}

	fmt.Println(x)

}
