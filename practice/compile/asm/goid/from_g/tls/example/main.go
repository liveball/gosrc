package main

import (
	"fmt"

	"gosrc/practice/compile/asm/goid/from_g/tls"
)

func main() {
	//tls.Setg(nil)
	fmt.Println("GetGoid:", tls.GetGoid())

	str := "addsdsdsds"
	fmt.Println("str len:", tls.StringLen(str))
	fmt.Println("str out:", tls.String(str))

	fmt.Println(1 & 1)

	stu := tls.Student{Name: "小韩", Age: 28}
	fmt.Println("stu:", tls.NewStudent(stu))

	nStu := tls.NewStudentPtr(&stu)
	fmt.Println("nStu1:", *nStu)
	nStu.Name = "小强"
	nStu.Age = 30
	fmt.Println("nStu2:", *nStu)

	tls.UpStudentPtr(nStu,"小王",40)
	fmt.Println("nStu up:", nStu)
}
