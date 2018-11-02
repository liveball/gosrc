package main

import (
	"fmt"
	"time"
)

//go build -o main -a -gcflags="-N -l -m"  main.go

// go tool objdump -S "main\.main" main > main.s

// GODEBUG=gctrace=1 $GODEV/bin/go run main.go

//Person for one people
type Person struct {
	Name string
	Age  int
}

func main() {
	// p := Person{}
	// modify(p)
	// fmt.Printf("person(%+v)\n", p)

	mynew()
}

func mynew() {
	rect1 := new(Person)
	rect1.Name = "xxx"
	rect1.Age = 22
	fmt.Printf("%v  %T  %v \n", rect1, rect1, *rect1)

	rect2 := &Person{"阿呆", 25}
	fmt.Printf("%v  %T  %v \n", rect2, rect2, *rect2)

	rect3 := Person{"小明", 26}
	fmt.Printf("%v  %T\n", rect3, rect3)
}

func modify(p Person) {
	p.Name = "tom"
}

func empty() {
	x := struct {
	}{}

	// x := new(struct { // x size ?
	// 	_ struct{}
	// })

	// x := &struct {
	// }{}

	_ = x

	// go func() {
	// 	for {
	// 		var m runtime.MemStats
	// 		runtime.ReadMemStats(&m)

	// 		log.Printf(
	// 			"Alloc(%v) TotalAlloc(%v) Sys(%v) Lookups(%v) Mallocs(%v) Frees(%v)\n",
	// 			float64(m.Alloc)/1024/1024,
	// 			float64(m.TotalAlloc)/1024/1024,
	// 			float64(m.Sys)/1024/1024,
	// 			float64(m.Lookups)/1024/1024,
	// 			float64(m.Mallocs)/1024/1024,
	// 			float64(m.Frees)/1024/1024,
	// 		)

	// 		time.Sleep(3 * time.Second)
	// 	}
	// }()

	// fmt.Println("hello")
	time.Sleep(100 * time.Second)

	// s := make([]struct{}, 5)
	// _ = s

	// fmt.Printf("s len(%d) s[0](%p) s[4](%p)\n", len(s), &s[0], &s[4])
	// fmt.Printf("s size(%d)  value(%#v)\n",
	// 	unsafe.Sizeof(s),
	// 	(*reflect.SliceHeader)(unsafe.Pointer(&s)),
	// )
}
