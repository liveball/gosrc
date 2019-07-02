package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	na := slice(a, 3, 6).([]int)
	fmt.Println(na)

	b := []string{"a", "b", "c", "d", "e", "f"}
	nb := slice(b, 2, 4).([]string)
	fmt.Println(nb)
}

func slice(l interface{}, s, e int) interface{} {
	rl := reflect.ValueOf(l)
	cnt := rl.Len() - 1

	switch {
	case s > cnt || s > e:
		return rl.Slice(0, 0).Interface()
	case e >= cnt:
		return rl.Slice(s-1, cnt).Interface()
	default:
		return rl.Slice(s-1, e).Interface()
	}
}
