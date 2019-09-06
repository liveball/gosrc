// +build ignore

package main

import . "github.com/mmcloughlin/avo/build"

func main() {
	TEXT("swapInt32", NOSPLIT, "func(addr *int32, newer int32) (old int32)")
	Doc("swapInt32 adds addr and newer.")
	addr := Load(Param("addr"), GP64())
	newer := Load(Param("newer"), GP64())
	ADDQ(addr, newer)
	Store(newer, ReturnIndex(0))
	RET()
	Generate()
}
