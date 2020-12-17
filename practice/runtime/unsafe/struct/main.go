package main

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"text/template"
	"time"
	"unsafe"
)

type user struct {
	name string
	age  int
}

// MyTemplate 定义和 template.Template 只是形似
type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}

// Template is the representation of a parsed template. The *parse.Tree
// field is exported only for use by html/template and should be treated
// as unexported by all other clients.
//type Template struct {
//	name string
//	*parse.Tree
//	*common
//	leftDelim  string
//	rightDelim string
//}

type a struct {
	foo string
	bar int
}

type b struct {
	foo string
	bar int
}

func main() {
	var aa *a
	aa = new(a)
	bb := (*b)(unsafe.Pointer(&aa))
	_ = bb.foo
	fmt.Println(111, bb)
	//fmt.Println(bb.foo)//内存暴增。。。

	t := template.New("Foo")
	p := (*MyTemplate)(unsafe.Pointer(t))
	p.name = "Bar" // 关键在这里，突破私有成员
	fmt.Println(p, t)

	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张三"

	// pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	// *pAge = 20

	//error
	// 逻辑上看，以上代码不会有什么问题，但是这里会牵涉到GC，如果我们的这些临时变量被GC，那么导致的内存操作就错了，
	// 我们最终操作的，就不知道是哪块内存了，会引起莫名其妙的问题
	// temp := uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)
	// pAge := (*int)(unsafe.Pointer(temp))
	// *pAge = 30

	fmt.Println(*u)
	// ref: https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html

	// emptyStruct()
}

func emptyStruct() {
	x := struct {
	}{}

	// x := new(struct { // x size ?
	// 	_ struct{}
	// })

	// x := &struct {
	// }{}

	_ = x

	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			log.Printf(
				"Alloc(%v) TotalAlloc(%v) Sys(%v) Lookups(%v) Mallocs(%v) Frees(%v)\n",
				float64(m.Alloc)/1024/1024,
				float64(m.TotalAlloc)/1024/1024,
				float64(m.Sys)/1024/1024,
				float64(m.Lookups)/1024/1024,
				float64(m.Mallocs)/1024/1024,
				float64(m.Frees)/1024/1024,
			)

			time.Sleep(3 * time.Second)
		}
	}()

	fmt.Println("hello")
	time.Sleep(100 * time.Second)
}

func sliceStruct() {
	s := make([]struct{}, 5)
	_ = s

	fmt.Printf("s len(%d) s[0](%p) s[4](%p)\n", len(s), &s[0], &s[4])
	fmt.Printf("s size(%d)  value(%#v)\n",
		unsafe.Sizeof(s),
		(*reflect.SliceHeader)(unsafe.Pointer(&s)),
	)
}
