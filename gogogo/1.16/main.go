package main

import (
	"embed"
	"fmt"
)

//go:embed *.txt
var global embed.FS

func main(){
	testFiles(global, "a.txt", "Concurrency is not parallelism.\n")
	testFiles(global, "b.txt", "hello, world\n")
	testFiles(global, "c.txt", "I can eat glass and it doesn't hurt me.\n")
}

func testFiles(f embed.FS, name, data string) {
	d, err := f.ReadFile(name)
	if err != nil {
		fmt.Printf("f.ReadFile name(%v) error(%v)\n", name, err)
		return
	}
	if string(d) != data {
		fmt.Printf("read %v = %q, want %q \n", name, d, data)
	}
}
