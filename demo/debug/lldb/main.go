package main

import "fmt"

// go build -gcflags "-N -l" -o main

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
