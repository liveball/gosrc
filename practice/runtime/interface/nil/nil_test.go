package main

import (
	"fmt"
	"testing"
)

type Student struct {
	Name string
}

func Test_1(t *testing.T) {
	var b interface{} = Student{
		Name: "aaa",
	}

	var c = b.(Student)
	c.Name = "bbb"

	fmt.Println(b.(Student).Name)
}

func Test_2(t *testing.T) {
	a := Student{Name: "aaa"}
	var b interface{} = a

	a.Name = "bbb"
	fmt.Println(b.(Student).Name)
}

func Test_3(t *testing.T) {
	var b interface{} = &Student{
		Name: "aaa",
	}

	var c = b.(*Student)
	c.Name = "bbb"

	fmt.Println(b.(*Student).Name)
}

func Test_4(t *testing.T) {
	var a interface{} = nil
	fmt.Println(a == nil)

	//当一个指针赋值给 interface 类型时，无论此指针是否为 nil，赋值过的 interface 都不为 nil
	var b *Student = nil
	var c interface{} = b
	fmt.Println(c == nil)
}

type Talkable interface {
	TalkEnglish(string)
	TalkChinese(string)
}

type Student1 struct {
	Talkable
	Name string
	Age  int
}

func Test_5(t *testing.T) {
	a := &Student1{Name: "aaa", Age: 12}
	var b Talkable = a // 这行相当于绑定了结构的类型，并没有去绑定结构体对应的方法集？
	fmt.Println(b)
}

func Test_6(t *testing.T) {
	a := Student1{Name: "bbb", Age: 15}
	a.TalkEnglish("nice to meet you\n")
}

func (s *Student1) TalkEnglish(s1 string) {
	fmt.Printf("I'm %s,%d years old,%s", s.Name, s.Age, s1)
}

func (s *Student1) TalkChinese(s1 string) {
	fmt.Printf("我是 %s, 今年%d岁,%s", s.Name, s.Age, s1)
}
