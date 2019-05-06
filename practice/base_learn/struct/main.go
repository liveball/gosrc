package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
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

type demo struct {
	age int
}

func (d *demo) Test() string {
	fmt.Println("call Test()")
	return fmt.Sprintf("%v", d.age)
}

// func (d demo) String() string {
// 	fmt.Println("call String()")
// 	return fmt.Sprintf("%v", d.age)
// }

func (d *demo) String() string {
	fmt.Println("1")
	return fmt.Sprintf("%v", d.age)
}

func main() {
	// ExampleMethodSet()
	// p := Person{}
	// modify(p)
	// fmt.Printf("person(%+v)\n", p)
	// mynew()
	// obj := Person{Age: 111, Name: "ddsdsds"}
	// _ = reflect.TypeOf(obj)

	//方法接收者和type 接口实现
	d := demo{5}
	// d1 := &demo{5}
	_, _ = fmt.Println(d)
	// _, _ = fmt.Println(d1)

	input, err := os.OpenFile("./main.go", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", input, 0)
	if err != nil {
		log.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("main", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	demoType := pkg.Scope().Lookup("demo").Type()
	fmt.Println("demoType:", demoType)
	for _, t := range []types.Type{demoType, types.NewPointer(demoType)} {
		fmt.Printf("Method set of %s:\n", t)
		mset := types.NewMethodSet(t)
		for i := 0; i < mset.Len(); i++ {
			fmt.Println(mset.At(i))
		}
		fmt.Println()
	}

}

// ExampleMethodSet prints the method sets of various types.
func ExampleMethodSet() {
	// Parse a single source file.
	const input = `
package temperature
import "fmt"
type Celsius float64
func (c Celsius) String() string  { return fmt.Sprintf("%g°C", c) }
func (c *Celsius) SetF(f float64) { *c = Celsius(f - 32 / 9 * 5) }
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "celsius.go", input, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Type-check a package consisting of this file.
	// Type information for the imported packages
	// comes from $GOROOT/pkg/$GOOS_$GOOARCH/fmt.a.
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("temperature", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print the method sets of Celsius and *Celsius.
	celsius := pkg.Scope().Lookup("Celsius").Type()
	fmt.Println(111, celsius)
	for _, t := range []types.Type{celsius, types.NewPointer(celsius)} {
		fmt.Printf("Method set of %s:\n", t)
		mset := types.NewMethodSet(t)
		for i := 0; i < mset.Len(); i++ {
			fmt.Println(mset.At(i))
		}
		fmt.Println()
	}

	// Output:
	// Method set of temperature.Celsius:
	// method (temperature.Celsius) String() string
	//
	// Method set of *temperature.Celsius:
	// method (*temperature.Celsius) SetF(f float64)
	// method (*temperature.Celsius) String() string
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
