package impl

import (
	_ "unsafe" // for go:linkname
)

//go:linkname helloWorld impl.helloWorld
func helloWorld() string {
	return "hello world"
}
