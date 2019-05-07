package main

// func add(a, b int) int {
// 	return a + b
// }

func add(a int) (r int) {
	b := 200
	b = a + b
	r = a + b
	return r
}

func main() {
	println(add(1))
}
