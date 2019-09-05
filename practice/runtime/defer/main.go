package main

type Test struct{}

func (p *Test) Stop() {
	println("defer1 stop")
}

type I interface {
	Stop()
}

func Start() I {
	println("defer1")
	var t Test
	return &t
}

func Start1() interface{ Stop() } {
	println("defer1")
	var t Test
	return &t
}

func Start2() {
	println("defer2")
}

func main() {
	//1.调用start 生成I的实例，再实现I的stop方法
	s := Start()
	defer s.Stop()

	//2.直接返回未命名 interface{ Stop() } 的接口，并且实现接口的stop方法
	defer Start1().Stop()

	defer Start2()
	println("main")

}
