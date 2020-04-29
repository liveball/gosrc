package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder
	sb.WriteString("aaaa")
	fmt.Println(sb.Len(), sb.Cap())
}
