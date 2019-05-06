package main

import (
	"fmt"
)

type Str string

func (s Str) String() string {
	// return fmt.Sprintf("Str: %s", s)//wrong
	return fmt.Sprintf("Str: %s", string(s)) //right
}

// You are implementing Str.String in terms of itself.
// return fmt.Sprintf("Str: %s", s) will call s.String(),
// resulting in infinite recursion. Convert s to string first.

// This is working as intended, you are using the %s verb to call Str's String method,
// which uses fmt.Sprint to call Str's String method, and so on.

func main() {
	var s Str = "hi"
	fmt.Println(s)
}
