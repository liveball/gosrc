package main

import (
	"sync"
)

const (
	cnt = 10000
)

var (
	batchSize = 100
	mids      = make([]int, 0, cnt)
)

func main() {
	batchG()
}

func init() {
	for i := 0; i < cnt; i++ {
		mids = append(mids, i)
	}
}

func batchG() {
	midsCount := len(mids)
	page := midsCount/batchSize + 1
	// pool := make([][]int, page)

	pool := make([][]int, 0, page)

	var wg sync.WaitGroup
	wg.Add(page)
	for i := 0; i < page; i++ {
		var (
			start = i * batchSize
			end   = (i + 1) * batchSize
			// currentPage = i
		)
		if i == page-1 {
			end = midsCount
		}

		go func() {
			tmp := mids[start:end]
			// pool[currentPage] = tmp
			pool = append(pool, tmp)
			wg.Done()
		}()
	}
	wg.Wait()

	// for _, v := range pool {
	// 	fmt.Println(v)
	// }
}
