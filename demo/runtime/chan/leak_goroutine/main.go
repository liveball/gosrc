package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

type achTarget struct {
	MID   int64 //用户mid
	Value int64 //指标类型
	Type  int32 //指标类型
}

var (
	chanSize                                      = 1024
	batchSize                                     = 10
	viewTargetChan, likeTargetChan, fanTargetChan chan *achTarget
	starTargetChan, finTaskTargetChan             chan *achTarget
)

func main() {
	viewTargetChan = make(chan *achTarget, chanSize)
	likeTargetChan = make(chan *achTarget, chanSize)
	fanTargetChan = make(chan *achTarget, chanSize)
	starTargetChan = make(chan *achTarget, chanSize)
	finTaskTargetChan = make(chan *achTarget, chanSize)

	stasticGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Println("NumGoroutine:", total)
		}
	}

	go stasticGroutine()

	var i1, i2, i3, i4, i5 int64
	go func() { //播放
		for {
			i1 += 1000
			viewTargetChan <- &achTarget{
				MID:   i1,
				Value: i1 % 100,
				Type:  int32(20),
			}
			time.Sleep(time.Second * 1)
		}
	}()

	go func() { //点赞
		for {
			i2 += 2000
			likeTargetChan <- &achTarget{
				MID:   i2,
				Value: i2 % 100,
				Type:  int32(20),
			}
			time.Sleep(time.Second * 2)
		}
	}()

	go func() { //粉丝
		for {
			i3 += 3000
			fanTargetChan <- &achTarget{
				MID:   i3,
				Value: i3 % 100,
				Type:  int32(20),
			}
			time.Sleep(time.Second * 3)
		}
	}()

	go func() { //新星计划
		for {
			i4 += 4000
			starTargetChan <- &achTarget{
				MID:   i4,
				Value: i4 % 100,
				Type:  int32(20),
			}
			time.Sleep(time.Second * 10)
		}
	}()

	go func() { //完成新手任务
		for {
			i5 += 5000
			finTaskTargetChan <- &achTarget{
				MID:   i5,
				Value: i5 % 100,
				Type:  int32(20),
			}
			time.Sleep(time.Second * 5)
		}
	}()
	mergeMsg()

	select {}
}
func mergeMsg() { //合并所有databus消息进行业务逻辑处理
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		defer ticker.Stop()
		res := make(map[int64]map[int32]int64)
		for {
			select {
			case v, ok := <-mergeTarget(
				viewTargetChan, likeTargetChan, fanTargetChan,
				starTargetChan, finTaskTargetChan):
				if !ok {
					return
				}

				tgMap := make(map[int32]int64)
				tgMap[v.Type] = v.Value
				res[v.MID] = tgMap

				println(111, v, len(res))
				if len(res) < batchSize {
					continue
				}

			case <-ticker.C:
				log.Println("mergeMsg s.mergeTarget wait chan data more than 60s")
			}

			if len(res) > 0 {
				log.Println(res)
				res = make(map[int64]map[int32]int64)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func mergeTarget(ats ...<-chan *achTarget) <-chan *achTarget {
	var wg sync.WaitGroup
	out := make(chan *achTarget, chanSize)
	output := func(c <-chan *achTarget) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(ats))
	for _, c := range ats {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
