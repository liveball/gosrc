package main

import (
	"fmt"
)

func main() {
	gcflags := []string{"-+", "-N", "-m", "-l"}
	for i := 0; i < len(gcflags); i++ {
		if gcflags[i] == "-N" { //i=1
			copy(gcflags[i:], gcflags[i+1:])
			gcflags = gcflags[:len(gcflags)-1] //gcflags [-+ -m -l -l]
			i--
		}
	}
	fmt.Println(gcflags)
}
