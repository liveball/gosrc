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

func exec() {
	for {
		t := time.Now()
		weekday := t.Weekday()
		if int(weekday) != 4 {
			continue
		}

		// todayZero := t.Truncate(time.Hour).Add(time.Duration(-t.Hour()) * time.Hour)
		// fmt.Printf("todayZero %+v\n", todayZero)

		startMtime := t.AddDate(0, 0, -7).Add(time.Hour * time.Duration(0)).Add(time.Minute * time.Duration(0)).Add(time.Second * time.Duration(10))
		endMtime := t.Add(time.Hour * time.Duration(0)).Add(time.Minute * time.Duration(0)).Add(time.Second * time.Duration(10))
		// sendMtime := t.Add(time.Hour * time.Duration(0))

		println(t.Hour(), t.Minute(), t.Second())
		sendMtime := (t.Unix() - int64(3600*t.Hour()) - int64(60*t.Minute()) - int64(t.Second())) + int64(20*3600) + int64(24*60)

		fmt.Printf("startMtime %+v\n", startMtime)
		fmt.Printf("endMtime %+v\n", endMtime)
		fmt.Printf("sendMtime %+v\n", sendMtime)
		println("----------------------")

		println(t.Unix(), sendMtime)
		if t.Unix() == sendMtime {
			println("aaaa")
		}

		time.Sleep(10 * time.Second)
	}
}

func intervalTimer(start time.Time, hour, minute, second int) <-chan time.Time {
	fmt.Printf("start %+v\n", start)
	next := start.Add(time.Hour * time.Duration(hour)).Add(time.Minute * time.Duration(minute)).Add(time.Second * time.Duration(second))
	fmt.Printf("next %+v\n", next)
	next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
	fmt.Printf("next.Sub(now) %+v\n", next.Sub(start))

	t := time.NewTimer(next.Sub(start)) //获取超时定时器
	return t.C
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
