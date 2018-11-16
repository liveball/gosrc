package main

import (
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			if !timeUp(10) {
				time.Sleep(time.Second)
				println("timeUp")
				continue
			}

			for {
				time.Sleep(time.Second)
				println("wait for exit")
			}
		}
	}()

	wg.Wait()

}

func timeUp(delay time.Duration) bool {
	now := time.Now()
	dayN := int(now.Weekday()) % 7
	timestamp := now.AddDate(0, 0, -dayN).Unix() - int64(now.Second()) - int64(60*now.Minute()) - int64(3600*now.Hour()) + int64(delay)
	return now.Unix() == timestamp
}
