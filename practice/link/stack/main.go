package main

func main() {
	// do not morestack
	// f(1, 2)

	// morestack main.f2
	f2(1, 2)
	// StackGuard= 1760 StackSystem= 0 StackSmall= 128
	// StackLimit = StackGuard - StackSystem - StackSmall= 1632
	// main.f2 top(function at top of safe zone once) true 8 objabi.StackLimit-callsize(ctxt)= 1624
}

func f(a, b int) (int, int) {
	sum := 0

	sum += a
	sum += b

	return sum, a + b
}

// 根据调试需要设置 //go:nosplit (禁止堆栈溢出检查 标记)
func f2(a, b int) (int, int) {
	sum := 0

	elems := make([]int, 100)
	for _, v := range elems {
		sum += v
	}

	sum += a
	sum += b

	return sum, a + b
}
