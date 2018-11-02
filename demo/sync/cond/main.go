package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func test(x int) {
	cond.L.Lock() //获取锁
	cond.Wait()   //等待通知 暂时阻塞
	fmt.Println(x)
	time.Sleep(time.Second * 1)
	cond.L.Unlock() //释放锁
}
func main() {
	for i := 0; i < 40; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 3)

	fmt.Println("signal1")
	cond.Signal() // 下发一个通知随机给已经获取锁的goroutine
	time.Sleep(time.Second * 3)

	fmt.Println("signal2")
	cond.Signal() // 下发第二个通知随机给已经获取锁的goroutine

	time.Sleep(time.Second * 1) // 在广播之前要等一会，让所有线程都在wait状态
	fmt.Println("broadcast")
	cond.Broadcast() //下发广播给所有等待的goroutine

	time.Sleep(time.Second * 60)
}
