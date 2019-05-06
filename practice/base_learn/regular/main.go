package main

import (
	"fmt"
	"regexp"
)

func main() {
	qq := regexp.MustCompile(`^\d{2,4}$`)
	fmt.Println(qq.MatchString("789"))
	fmt.Println(qq.MatchString("789111"))

	matched, err := regexp.MatchString("foo.*", "seafood")
	fmt.Println(matched, err)

	matched2, err := regexp.MatchString("sea.*", "seafood")
	fmt.Println(matched2, err)
}
