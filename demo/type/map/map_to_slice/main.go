package main

import "fmt"

func main() {
	midMap := make(map[int]int)
	for i := 1; i <= 50; i++ {
		midMap[i] = i
	}

	mids := make([]int, 0, len(midMap))
	for mid := range midMap {
		mids = append(mids, mid)
	}

	var tmids []int
	batchSize := 10
	count := len(mids)/batchSize + 1
	println(count)
	for i := 0; i < count; i++ {
		if i == count-1 {
			println(i)
			tmids = mids[i*batchSize:] //最后一批
		} else {
			println(i, (i + 1))
			tmids = mids[i*batchSize : (i+1)*batchSize] //i 到 i+1批
		}

		if len(tmids) > 0 {
			fmt.Println(tmids)
		}
	}
}
