package main

import "fmt"

type Talkable interface {
	TalkEnglish(string)
	TalkChinese(string)
}

func main() {
	noBindMethods() //不绑定方法集

	//bindMethods() //绑定方法
}

type Student struct {
	Talkable
	Name string
	Age  int
}

func bindMethods() {
	a := &Student{Name: "aaa", Age: 12}
	var b Talkable = a //

	a.TalkEnglish("nice to meet you\n")

	fmt.Println(b)
}

func (s *Student) TalkEnglish(s1 string) {
	fmt.Printf("I'm %s,%d years old,%s", s.Name, s.Age, s1)
}

func (s *Student) TalkChinese(s1 string) {
	fmt.Printf("我是 %s, 今年%d岁,%s", s.Name, s.Age, s1)
}

type Student1 struct {
	Talkable
	Name string
	Age  int
}

func noBindMethods() {
	a := &Student1{Name: "aaa", Age: 12}
	var b Talkable = a // 这行相当于绑定了结构的类型，并没有去绑定结构体对应的方法集？

	fmt.Println(b)
}
