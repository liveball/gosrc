package main

type ISuppliers interface {
	GetName() string
	Send(a int) bool
}

type Test struct{}

func (t *Test) GetName() string {
	return "test"
}

func (t *Test) Send(a int) bool {
	return true
}

func GetList() []ISuppliers {
	var is []ISuppliers
	is = append(is, &Test{})
	return is
}

func main() {
	is := GetList()

	for _, v := range is {
		println("v.GetName",v.GetName())
	}
}
