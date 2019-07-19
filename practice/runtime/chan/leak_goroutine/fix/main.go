package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"net/http"
	_ "net/http/pprof"
)

//浏览器中查看 http://localhost:9000/debug/pprof/
//go tool pprof http://localhost:9000/debug/pprof/profile

type achTarget struct {
	MID   int64
	Value int64
	Type  int32
}

var (
	chanSize          = 1024
	batchSize         = 100
	viewTargetChan    = make(chan *achTarget, chanSize)
	likeTargetChan    = make(chan *achTarget, chanSize)
	fanTargetChan     = make(chan *achTarget, chanSize)
	starTargetChan    = make(chan *achTarget, chanSize)
	finTaskTargetChan = make(chan *achTarget, chanSize)
)

func main() {
	mock()
	mergeMsg()

	stasticGroutine := func() {
		for {
			time.Sleep(time.Second)
			total := runtime.NumGoroutine()
			fmt.Println("NumGoroutine:", total)
		}
	}
	go stasticGroutine()

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func mergeMsg() {
	//处理方式1  for { select{}}
	go func() {
		ticker := time.NewTicker(time.Second * 60)
		defer ticker.Stop()
		res := make(map[int64]map[int32]int64)
		out := mergeTarget(viewTargetChan, likeTargetChan, fanTargetChan, starTargetChan, finTaskTargetChan)
		for {
			select {
			case v, ok := <-out:
				if !ok {
					return
				}

				tgMap := make(map[int32]int64)
				tgMap[v.Type] = v.Value
				res[v.MID] = tgMap

				if len(res) < batchSize {
					continue
				}

			case <-ticker.C:
				log.Println("wait chan data more than 60s")
			}

			if len(res) > 0 {
				//log.Println(len(res), res)
				res = make(map[int64]map[int32]int64)
				time.Sleep(10 * time.Millisecond)
			}
		}

	}()

	//处理方式2
	// go func() {
	// for v := range mergeTarget(viewTargetChan, likeTargetChan, fanTargetChan, starTargetChan, finTaskTargetChan) {

	// 	tgMap := make(map[int32]int64)
	// 	tgMap[v.Type] = v.Value
	// 	res[v.MID] = tgMap

	// 	println(111, v, len(res))
	// 	if len(res) < batchSize {
	// 		continue
	// 	}

	// 	if len(res) > 0 {
	// 		log.Println(res)
	// 		res = make(map[int64]map[int32]int64)
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }
	// }()
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

func mock() {
	var i1, i2, i3, i4, i5 int64
	go func() {
		for {
			i1 += 11
			viewTargetChan <- &achTarget{
				MID:   i1,
				Value: i1,
				Type:  int32(10),
			}
		}
	}()

	go func() {
		for {
			i2 += 2
			likeTargetChan <- &achTarget{
				MID:   i2,
				Value: i2,
				Type:  int32(20),
			}
		}
	}()

	go func() {
		for {
			i3 += 3
			fanTargetChan <- &achTarget{
				MID:   i3,
				Value: i3,
				Type:  int32(30),
			}
		}
	}()

	go func() {
		for {
			i4 += 4
			starTargetChan <- &achTarget{
				MID:   i4,
				Value: i4,
				Type:  int32(40),
			}
		}
	}()

	go func() {
		for {
			i5 += 5
			finTaskTargetChan <- &achTarget{
				MID:   i5,
				Value: i5,
				Type:  int32(50),
			}
		}
	}()
}
