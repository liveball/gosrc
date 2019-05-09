package main

//Interface def
type Interface interface {
	Forward(a, b int) int
	Reverse(a, b int) int
}

//I iii
type I struct {
	Interface
}

//Decompost 分解
func (i *I) Decompost(a, b int) int {
	return i.Forward(a, b)
}

func main() {
	i := new(I)
	i.Decompost(1, 2)
}
