package main

func main() {
	_ = add(1, 2)
}

func add(a, b int) int {
	foo()
	return a + b
}


func foo() int{
	return 100
}