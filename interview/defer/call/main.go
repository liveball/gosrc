package main

func main() {
	call()
}

func call() {
	defer func() { println("打印前") }()
	defer func() { println("打印中") }()
	panic("panic")
	defer func() { println("打印后") }()
}
