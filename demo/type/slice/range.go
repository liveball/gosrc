package main

import "fmt"

func main() {
	a := [3]int{0, 1, 2}
	for k, v := range a {
		if k == 0 {
			a[1], a[2] = 998, 999
			fmt.Println(a)
		}

		a[k] = v + 100
		vv := v
		fmt.Println(&v, vv, &vv)
	}
	fmt.Println(a)

	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 6
		}
		vv := v
		fmt.Println(i, v, &v, vv, &vv)
	}
	fmt.Println(s)
}
