package main

import (
	"fmt"
	"time"
)

func main() {
	dispatchTask()

	for i:=0;i<100;i++{
		checkTaskChan<-int64(i)
	}

	time.Sleep(time.Second*10)
	taskDone<- struct{}{}
}

var (
	checkTaskChan = make(chan int64, 100)
	taskDone      = make(chan struct{})
	taskQueue = make([]int64, 0, 1024)
)

func dispatchTask() {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)
		defer ticker.Stop()
		for {
			select {
			case id := <-checkTaskChan://从chan 里面消费
				taskQueue = append(taskQueue, id)
			case <-ticker.C://定时从切片获取处理
				if len(taskQueue) > 0 {
					id := taskQueue[0]
					taskQueue = taskQueue[1:]
					fmt.Println(111, id)
				}
			case <-taskDone:
				taskQueue = taskQueue[:0]//退出的时候清空
				fmt.Println("CheckTask close")
				return
			}
		}
	}()
}
