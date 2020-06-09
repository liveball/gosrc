package main

import "fmt"

func main() {
	b := []int{1, 2, 3, 4, 5}

	bb := make(map[int]*int)
	for k, v := range b {

		bb[k] = &v //55555
		//bb[k] = &b[k]
	}

	for _, v := range bb {
		fmt.Println(*v)
	}
}
