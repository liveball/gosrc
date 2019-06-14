package main

import (
	_ "unsafe"
)

func SwapInt32(addr *int32, new int32) (old int32)

func main() {
	i := int32(10)
	j := int32(20)
	SwapInt32(&i, j)
}
