package main


func main(){
	//yes
	a:=Fetch2("a")
	b:=Fetch2("b")

	//consume(<-a, <-b)

	//no
	//a:=<-Fetch2("a")
	//b:=<-Fetch2("b")
	//
	////consume(a, b)

//	for item := range Glob("[ab]*") {
//		[...]
//	}

}

func Fetch(name string, f func(Item)) {
	go func() {
		f(item)
	}()
}

func Fetch2(name string) <-chan Item{
	c:=make(chan Item, 1)
	go func() {
		//[...]
		c<-item
	}()

	return c
}


func Glob(pattern string) <-chan Item{
	c:=make(chan Item)
	go func() {
		defer close(c)

		//for [...]{
		//	[...]
		//	c<-item
		//}
	}()

	return c
}


func Async(x In)(<-chan Out){
	c:=make(chan Out, 1)
	go func() {
		c<-Synchronous(x)
	}()
	
	return c
}

func Synchronous(in In) Out {
	c:=Async(x)
	return <-c
}
