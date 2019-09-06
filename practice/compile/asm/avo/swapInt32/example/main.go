package main

import "gosrc/practice/compile/asm/swapInt32"

func main() {
	i := int32(10)
	j := int32(20)
	swapInt32.SwapInt32(&i, j)
}