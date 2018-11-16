package main

import (
	"fmt"
	"math"
	"time"
)

var (
	count             = 3
	maxConsume        = 100
	maxJob            = 100
	rewardNotifyQueue = make([]chan int, maxConsume)
)

func main() {
	exec()
	select {}
}

func day() {
	year, month, day := time.Now().Date()
	stoday := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	etoday := time.Date(year, month, day, 23, 59, 59, 999, time.Local)
	syesday := stoday.AddDate(0, 0, -1)
	eyesday := etoday.AddDate(0, 0, -1)
	println(eyesday.Unix() - syesday.Unix())
}

func exec() {
	for {
		now := time.Now()
		if int(now.Weekday()) != 5 || now.Hour() != 10 { //|| now.Minute() != 21
			time.Sleep(1 * time.Second)
			println("每周星期4 23点15分运行一次")
			continue
		}

		fmt.Printf("now %+v\n", now)
		last := now.AddDate(0, 0, -7).Add(time.Minute * time.Duration(1))
		fmt.Printf("last %+v\n", last)

		<-intervalTimer(now, 0, 0, 10).C
		println("----------------------")

		time.Sleep(10 * time.Second)
	}
}

func intervalTimer(now time.Time, hour, minute, second int) *time.Timer {
	next := now.Add(time.Hour * time.Duration(hour)).Add(time.Minute * time.Duration(minute)).Add(time.Second * time.Duration(second))
	// next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Hour(), 0, 0, next.Location())
	fmt.Println("next", next.Format("2006-01-02 15:04:05"))
	return time.NewTimer(next.Sub(now))
}

func getWeekdaysBetween(start, end time.Time) int {
	offset := -int(start.Weekday())
	start = start.AddDate(0, 0, -int(start.Weekday()))

	offset += int(end.Weekday())
	if end.Weekday() == time.Sunday {
		offset++
	}
	end = end.AddDate(0, 0, -int(end.Weekday()))

	dif := end.Sub(start).Truncate(time.Hour * 24)
	weeks := float64((dif.Hours() / 24) / 7)
	return int(math.Round(weeks)*5) + offset
}

func getWeekday(now time.Time, dayOfWeek, hour, minute, second int) time.Time {
	h := time.Duration(-now.Hour()) * time.Hour
	t := now.Truncate(time.Hour).Add(h)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+dayOfWeek) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d).Add(time.Hour * time.Duration(hour)).Add(time.Duration(minute) * time.Minute).Add(time.Duration(second) * time.Second)
}

func ticker() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
