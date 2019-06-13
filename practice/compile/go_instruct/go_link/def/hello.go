package def

import (
	_ "unsafe" // for go:linkname

	_ "gosrc/practice/compile/go_instruct/go_link/impl"
)

//go:linkname helloWorld impl.helloWorld
func helloWorld() string

func Hello() {
	println(helloWorld())
}
