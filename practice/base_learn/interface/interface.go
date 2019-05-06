package main

import (
	"fmt"
	"reflect"
)

// underlying type （基础类型）

// named type  （命名类型）

type mPtr *int

func underlying() {

	type a string
	var b string = "qq"
	// var c a = b //cannot use b (type string) as type a in assignment
	fmt.Printf(" b(%+v) \n", reflect.TypeOf(b))

	var p *int       //p    unnamed type
	var ptr mPtr = p //ptr  named type
	fmt.Printf("p(%+v) ptr(%+v) \n", reflect.TypeOf(p), reflect.TypeOf(ptr))

	var m myMap = make(map[int]string) //m  named type
	var mm map[int]string = m          //mm unnamed type
	m.add(2, "b")
	fmt.Println(m)
	fmt.Printf("m(%+v) mm(%+v) \n", reflect.TypeOf(m), reflect.TypeOf(mm))
}

//如果两个类型的基础类型相同，则具有下列特性:

//1、如果两个type 都是named type,彼此之间不能相互赋值
// 2.如果两个type 其中一个是 unnamed type，彼此之间可以相互赋值

// named type 和 unnamed type不同
// 当 named types 被作为一个 function 的 receiver 时，
// 它就拥有了自己的方法，unamed types 则不能，这是它们的重要区别。

type myMap map[int]string

func (m myMap) add(key int, value string) {
	m[key] = value
}

//有个一例外是是 pre-declare types 不能拥有自己的方法
// func (n int) name() { //cannot define new methods on non-local type int
// 	print(n)
// }

// type 的属性继承一：直接继承
// declared named type 不会从它的 underlying type 或 existing type 继承 method，但是会继承 field。
type person struct {
	name string
}

func (p *person) Speak() {
	fmt.Println("I am a person")
}

type student person //直接继承

func inherit() {
	var p person
	p.Speak()
	var s student
	fmt.Printf("s(%+v) p(%+v) \n", reflect.TypeOf(p), reflect.TypeOf(s))
	s.name = "jone"
	fmt.Println(s.name)
	// s.Speak()//s.Speak undefined (type student has no field or method Speak)

}

// The declared type does not inherit any methods bound to the existing type,
// but the method set of an interface type or of elements of a composite type remains unchanged:

// But declared named type 例外的情况之一：如果 existing type 是 interface，它的 method set 会被继承。

//I for test
type I interface {
	Talk()
}

// II existing type 是 I，I 是个接口，可以直接继承 I 的方法，II 等同于 I
type II I

//Person for test interface
type Person struct {
	name string
}

//Speak for test interface
func (p *Person) Speak() {
	fmt.Println("I am a person")
}

//Talk implement
func (p *Person) Talk() {
	fmt.Println("I am talking")
}

func inheritByInterface() {
	var p Person
	p.Speak()
	p.Talk()
	var i I
	i = &p
	i.Talk()
	var ii II
	ii = &p
	ii.Talk()

	fmt.Printf("i(%+v) ii(%+v) \n", reflect.TypeOf(i), reflect.TypeOf(ii))
}

// type 的属性继承二：type embedding

// M test
type M interface {
	Talk()
}

//Stu for test
type Stu struct {
	name string
}

//Speak for test
func (p *Stu) Speak() {
	fmt.Println("I am a student")
}

//Talk for test
func (p *Stu) Talk() {
	fmt.Println("I am talking")
}

//People for test stu
type People struct {
	Stu
}

func embedding() {
	var people People
	people.name = "people"
	people.Speak()
	people.Talk()

	fmt.Printf("people(%+v) \n", reflect.TypeOf(people))
}
