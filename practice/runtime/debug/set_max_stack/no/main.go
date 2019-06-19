package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go print(i)
	}

	time.Sleep(time.Second)
}

func print(i int) {
	fmt.Println(i)
}
