package main

import (
	"fmt"
	"go/types"
	"reflect"
	"time"
)

func main() {
	msg := "Starting main"
	fmt.Println(msg)
	bus := make(chan int)
	msg = "starting a gorountie"
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
	// var a interface{} = (*interface{})(nil)

	// fmt.Printf("a type: %#v \n\n value:  %#v \n",
	// 	reflect.TypeOf(a),
	// 	reflect.ValueOf(a),
	// )

	// if a == nil {
	// 	fmt.Println("a is nil")
	// } else {
	// 	fmt.Println("a is not nil")
	// }

}

//使用 -gcflags "-N -l" 参数关闭编译器代码优化
//go build -gcflags "-N -l" -o main main.go

func dumpObj() {
	for _, name := range types.Universe.Names() {
		// println(name)
		if obj, _ := types.Universe.Lookup(name).(*types.TypeName); obj != nil {
			fmt.Printf("obj(%#v) \n", reflect.ValueOf(obj))
			println()
		}
	}
}

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		c <- i
	}
	close(c)
}
