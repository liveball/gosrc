package main

type foo struct{}

func (f *foo) add() {}

func (f foo) sub() {}

func main() {
	//结构体值类型，只能调用值接受者方法
	// foo{}.add() //静态编译出错，cannot call pointer method on foo literal, cannot take the address of foo literal
	foo{}.sub()

	//值类型变量，如果声明一个变量，则既可以调用值接受者方法，又可以调用指针接受者方法
	var fv foo
	fv.add()
	fv.sub()

	//指针类型变量，既可以调用值接受者方法，又可以调用指针接受者方法
	var fp *foo
	//不初始化指针静态编译没问题，运行时panic: runtime error: invalid memory address or nil pointer dereference
	fp = new(foo) //指针初始化方法1
	// fp = &foo{}//指针初始化方法2
	fp.add()
	fp.sub()
}
