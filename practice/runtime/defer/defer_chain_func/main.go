package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(i int) *Slice {
	*s = append(*s, i)

	fmt.Println(i)

	return s
}

func main() {
	s := NewSlice()

	defer s.Add(1).Add(2).Add(3) //defer 只把最外层函数压入栈

	s.Add(4)

	//fmt.Println(s) [1 2 4]
}
