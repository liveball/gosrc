package main

func foo() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return
}

func main() {
	println(foo())
}
