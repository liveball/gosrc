package main

func myfalse() bool {
	return false
}

func main() {
	switch myfalse() {
	case true:
		println("true")
	case false:
		println("false")
	}
}
