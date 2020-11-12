package main

import (
	"fmt"
	"strings"
)

func main() {

	etag := "abc"
	//etag = strings.TrimPrefix(etag, "a")
	//fmt.Println(etag)
	etag = strings.TrimSuffix(etag, "a")
	fmt.Println(etag)
}
