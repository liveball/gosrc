package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		buf := make([]byte, 64*1024)
	// 		buf = buf[:runtime.Stack(buf, false)]
	// 		pl := fmt.Sprintf("panic in cache proc, err: %v, stack: %s", r, string(buf))
	// 		fmt.Fprintf(os.Stderr, pl)
	// 	}
	// }()

	c1 := make(chan int, 10)
	fmt.Println(<-c1)
}
