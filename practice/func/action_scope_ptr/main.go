package main

func main() {
	var a *int
	a = new(int)
	*a = 1
	incr(a)
	println(*a)
}

func incr(i *int) {
	*i++
}
