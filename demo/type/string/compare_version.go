package main

import (
	"errors"
	"strings"
)

func main() {
	println(compare("0.0.11", "0.0.11"))

	println(compare("1.0", "0.1"))
	println(compare("0.1", "0.0.1"))
	println(compare("0.11.0", "0.12.0"))

	println(compare("1.3.4a", "0.2.4b"))
	println(compare("0.3.4a", "1.2.4a"))

	println(compare("1.3.4a", "1.2.4b"))
	println(compare("1.2.4b", "1.3.4a"))

	println(compare("1.2.3a", "1.2.4b"))
	println(compare("1.2.3a", "1.2.4b"))
}

// v1=1.2.3a v2=1.2.4b
func compare(v1, v2 string) (res string) {
	v1Sli := strings.Split(v1, ".")
	v2Sli := strings.Split(v2, ".")

	if len(v1Sli) == 0 && len(v2Sli) == 0 {
		return errors.New("invalid version").Error()
	}
	for k, v := range v1Sli {
		if v == v2Sli[k] {
			continue
		} else if v > v2Sli[k] {
			res = v1 + ">" + v2
			return
		} else if v < v2Sli[k] {
			res = v1 + "<" + v2
			return
		}
	}
	res = v1 + "=" + v2
	return res
}
