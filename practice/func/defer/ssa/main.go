package main

func f() (x int) {
	for {
		defer func() {
			recover()
			x = 1
			// println(222, x)
		}()
		// println(111, x)
		panic(1)
	}
}

func main() {
	println(f())
}
