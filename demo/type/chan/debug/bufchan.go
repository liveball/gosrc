package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	//outCh buf大小必须和generator 长度一致，因为inCh为无缓冲chan，outCh在所有goroutine执行完毕才读取数据
	//所以为了不阻塞inCh 必须设置outCh的缓冲为inCh 可读出的数据长度
	inCh := generator(10)
	outCh := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go do(&wg, inCh, outCh)
	}
	wg.Wait()
	close(outCh)

	for o := range outCh {
		fmt.Println(o)
	}

	// for {
	// 	select {
	// 	case o, ok := <-outCh:
	// 		if !ok {
	// 			return
	// 		}
	// 		fmt.Println(o)
	// 	}
	// }
}

func generator(n int) <-chan int {
	gen := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			gen <- i
		}
		close(gen)
	}()
	return gen
}

func do(wg *sync.WaitGroup, in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	wg.Done()
}
