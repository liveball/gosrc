package main

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
)

type Node struct {
	sync.Mutex
	// *sync.Mutex
	Data map[string]string
}

var Cache []Node

// var Cache []*Node

func main() {
	Cache = make([]Node, 2)
	//copy 导致使用了多个不同的锁
	Cache[0] = Node{Data: make(map[string]string)}
	Cache[1] = Node{Data: make(map[string]string)}

	//多份数据copy使用同一个锁
	// Cache[0] = Node{Data: make(map[string]string), Mutex: &sync.Mutex{}}
	// Cache[1] = Node{Data: make(map[string]string), Mutex: &sync.Mutex{}}

	typ := reflect.TypeOf(Cache[0])
	// fmt.Println(typ)
	y := reflect.New(typ).Elem()
	for i := 0; i < typ.NumField(); i++ {
		// 根据每个struct field的type 设置其值
		fieldT := typ.Field(i)
		fieldV := y.Field(i)
		kind := fieldT.Type.Kind()
		println(kind.String(), fieldV.Addr().String())
	}

	//使用指针，修改同一值
	// Cache = make([]*Node, 2)
	// Cache[0] = &Node{Data: make(map[string]string)}
	// Cache[1] = &Node{Data: make(map[string]string)}

	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			j := index % 2
			node := Cache[j]

			node.Lock()
			defer node.Unlock()
			node.Data[strconv.Itoa(index)] = strconv.Itoa(index)
		}(i)
	}
	wg.Wait()
	fmt.Println(Cache[0])
}
