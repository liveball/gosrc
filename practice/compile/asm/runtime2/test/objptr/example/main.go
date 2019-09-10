package main

import (
	"fmt"

	"gosrc/practice/compile/asm/runtime2/test/objptr"
)

func main() {

	foo()
	fmt.Println("===============")
	stu()
}

func foo() {
	objptr.HelloWorld()
	fmt.Println("\n")
	fmt.Println("neg:", objptr.Neg(10))
}


func stu(){
	str := "addsdsdsds"
	fmt.Println("str len:", objptr.StringLen(str))
	fmt.Println("str out:", objptr.String(str))

	fmt.Println("stu:", objptr.NewStudent(objptr.Student{Name: "小李", Age: 25}))

	stu := objptr.Student{Name: "小韩", Age: 28}

	nStu := objptr.NewStudentPtr(&stu)
	fmt.Println("nStu1:", *nStu)
	nStu.Name = "小强"
	nStu.Age = 30
	fmt.Println("nStu2:", *nStu)

	objptr.UpStudentPtr(nStu, "小王", 40)
	fmt.Println("nStu up:", nStu)

	addr := objptr.StudentPtr(nStu)
	mStu := (*objptr.Student)(addr)
	fmt.Println("ptr name:", mStu.Name)

	mStu.Name = "大王ptr"
	mStu.Age = 98
	fmt.Println("name:", objptr.StudentAge(mStu))

	getName := objptr.StudentName(mStu)
	fmt.Println("name:", getName)
}