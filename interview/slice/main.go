package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	students()
}

func students() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "张三", Age: 30},
		{Name: "李四", Age: 25},
		{Name: "王五", Age: 42},
	}

	for _, stu := range stus {
		fmt.Printf("stu:%+v, ptr:%p\n", stu, &stu)
		m[stu.Name] = &stu //wrong

		//fix
		// st := stu
		// m[stu.Name] = &st
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.Age)
	}
}
