package main

import (
	"fmt"
	"strconv"
)

var (
	testHookParse = func(
		fn func(string) (int, error),
		a string,
	) (int, error) {
		b, err := fn(a)
		fmt.Printf("b type(%T)\n", b)
		return b, err
	}
)

func parseFunc(s string) (int, error) {
	return strconv.Atoi(s)
}

func main() {
	testHookParse(parseFunc, "10")
}
