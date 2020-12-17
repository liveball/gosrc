package main

import (
	"fmt"
	"gosrc/go/src/encoding/json"
	"gosrc/go/src/log"
	"time"
)

var (
	//ct = make(map[string]interface{})

	ct map[string]interface{}
)

func main() {
	a := 0

	go func() {
		for {
			a++
			//time.Sleep(time.Second*1)
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println(a)

	//ct["a"] = 1
	str, err := json.Marshal(ct)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("json", string(str))
}

//1、逃逸分析
//2、go程cpu密集，未出让cpu，协程每次去主协程获取寄存器中存储的a内存地址
//3、for循环里加个sleep或者print之类的io操作，应该就不会被优化成取cache了
