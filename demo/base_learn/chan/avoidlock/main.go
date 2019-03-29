package main

import (
	"fmt"
	"sort"
	"sync"
)

var num = 100

func main() {
	// useChan()

	useLock()
}

func useChan() {
	var (
		wg        sync.WaitGroup
		countArr  = make([]int, 0, num)
		countChan = make(chan int, num)
	)

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			countChan <- getCount(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	close(countChan)

	for v := range countChan {
		countArr = append(countArr, v)
	}

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i] < countArr[j]
	})

	fmt.Println(countArr)
}

func useLock() {
	var (
		wg       sync.WaitGroup
		lock     sync.Mutex
		countArr = make([]int, 0, num)
	)

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			lock.Lock()
			countArr = append(countArr, getCount(i))
			lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()

	sort.Slice(countArr, func(i, j int) bool {
		return countArr[i] < countArr[j]
	})

	fmt.Println(countArr)
}

func getCount(i int) int {
	return i + 100
}
