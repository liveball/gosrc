package main

import (
	"fmt"
	"math/rand"
	"time"
)

// wait 实现wait方法，收集jobs中channel中的int值返回。
// 每个channel都一定有一个int返回，不会close。
// 返回slice长度应当与jobs相同，顺序任意。
func wait(jobs []<-chan int) []int {
	//panic("implement me")

	res := []int{}
	for _, job := range jobs {
		select {
		case i, ok := <-job:
			if !ok {
				return res
			}
			res = append(res, i)
			//default:
			//log.Println("xxx")
		}
	}

	return res

}

// 以下代码仅为示例、调试用，不用修改
func main() {
	rand.Seed(time.Now().Unix())
	jobs := begin()
	results := wait(jobs)
	/*预期输出例子如下
	produce 2 821
	produce 4 931
	produce 0 1023
	produce 1 1063
	produce 3 1188
	wait[1023 1063 821 1188 931]
	*/
	fmt.Print("wait", results)
}

func begin() (rst []<-chan int) {
	for i := 0; i < rand.Intn(10)+3; i++ {
		ch := make(chan int)
		ii := i
		go func() {
			// 模拟一些异步任务
			x := rand.Intn(1000) + 200
			time.Sleep(time.Millisecond * time.Duration(x))
			println("produce", ii, x)
			ch <- x
		}()
		rst = append(rst, ch)
	}

	return
}
