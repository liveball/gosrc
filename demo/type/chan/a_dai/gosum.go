package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//c := make(chan int, 1)
	m := make(chan map[string]interface{},2)

	go sum(1, 2, m)
	ret := <-m
	fmt.Println(ret)

	go dev("122", "456", m)

	//fmt.Println(<-c)
	ret = <-m
	fmt.Println(ret)
	fmt.Println(ret["token"])

	//some func or operation
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}

func dev(a, b string, c chan map[string]interface{}) {
	time.Sleep(5 * time.Second)
	m := make(map[string]interface{})
	m["token"] = a + b
	m["flag"] = true
	c <- m
}

func sum(x, y int, c chan map[string]interface{}) {
	time.Sleep(5 * time.Second)
	m := make(map[string]interface{})
	m["count"] = x + y
	c <- m
}
