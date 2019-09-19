package main

type foo struct {
}

func main() {
	var i interface{}

	var f foo
	i = f

	_ = i

}
