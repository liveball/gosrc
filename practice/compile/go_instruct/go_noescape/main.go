package main

var a = 1000

func main() {
	println(es2())
}

func es() *int {
	var i int
	return &i
}

// go tool compile "-m" -o main.o /data/app/go/src/gosrc/practice/compile/go_instruct/go_noescape/main.go

// /data/app/go/src/gosrc/practice/compile/go_instruct/go_noescape/main.go:3:6: can inline main
// /data/app/go/src/gosrc/practice/compile/go_instruct/go_noescape/main.go:7:6: can inline es
// /data/app/go/src/gosrc/practice/compile/go_instruct/go_noescape/main.go:9:9: &i escapes to heap
// /data/app/go/src/gosrc/practice/compile/go_instruct/go_noescape/main.go:8:6: moved to heap: i

//go:noescape 指令强制要求编译器将其分配到函数栈上
//go:noescape
func es2() *int
