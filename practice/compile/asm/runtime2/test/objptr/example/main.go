package main

import (
	"fmt"
	"gosrc/practice/compile/asm/runtime2/test/objptr"
)

func main() {
	str := "addsdsdsds"
	fmt.Println("str len:", objptr.StringLen(str))
	fmt.Println("str out:", objptr.String(str))

	fmt.Println(1 & 1)

	stu := objptr.Student{Name: "小韩", Age: 28}
	fmt.Println("stu:", objptr.NewStudent(stu))

	nStu := objptr.NewStudentPtr(&stu)
	fmt.Println("nStu1:", *nStu)
	nStu.Name = "小强"
	nStu.Age = 30
	fmt.Println("nStu2:", *nStu)

	objptr.UpStudentPtr(nStu, "小王", 40)
	fmt.Println("nStu up:", nStu)

	fmt.Println(objptr.StudentName(nStu))
}

