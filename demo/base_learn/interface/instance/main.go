package main

type P interface {
	GetUID() int64
}

type Person struct {
	UID int64
}

func (p Person) GetUID() int64 {
	return p.UID
}

func main() {
	var p P
	p = Person{111} //成功实现P接口
	println(p.GetUID())
}
