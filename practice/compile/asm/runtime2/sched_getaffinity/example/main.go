package main

import "gosrc/practice/compile/asm/sched_getaffinity"

func main()  {
	cpuCnt:=sched_getaffinity.GetProcCount()
	println(cpuCnt)
}
