package main

import (
	"fmt"
	"time"
)

func beginningOfDay(t time.Time) time.Time {
	// println(-t.Hour(), time.Duration(-t.Hour()), time.Hour)
	d := time.Duration(-t.Hour()) * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func getTuesday(now time.Time) time.Time {
	t := beginningOfDay(now)    //取当前时间开始的时刻
	weekday := int(t.Weekday()) //获取当前是星期几
	if weekday == 0 {
		weekday = 7
	}

	d := time.Duration(-weekday+6) * 24 * time.Hour //当前时间往前倒两天

	fmt.Printf("%v\n", t.Weekday())
	println(t.Month(), t.Day(), t.Weekday(), time.Duration(-weekday+2)*24)

	println(time.Unix(t.Truncate(time.Hour).Add(d).Unix(), 0).Format("2006-01-02 15:04:05"))
	return t.Truncate(time.Hour).Add(d)
}

func getSunday(now time.Time) time.Time {
	t := beginningOfDay(now)
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	}
	d := time.Duration(-weekday+7) * 24 * time.Hour //当前时间往前倒七天
	return t.Truncate(time.Hour).Add(d)
}

func getDate() (sd string) {
	t := time.Now()
	td := getTuesday(t).Add(12 * time.Hour)
	// if t.Before(td) { //当前时间在本周二12点之前，则取上上周日的数据，否则取上周日的数据
	// 	sd = getSunday(t.AddDate(0, 0, -14)).Format("20060102")
	// } else {
	// 	sd = getSunday(t.AddDate(0, 0, -7)).Format("20060102")
	// }
	fmt.Printf("current time (%s) tuesday (%s) sunday (%s)", t.Format("2006-01-02 15:04:05"), td, sd)
	return
}

func main() {
	getDate()
	now := time.Now()

	delay := 1

	dayN := int(now.Unix()/86400+4) % 7

	println(dayN)
	timestamp := now.AddDate(0, 0, -dayN).Unix() - int64(now.Second()) - int64(60*now.Minute()) - int64(3600*now.Hour()) + int64(delay)
	fmt.Println(time.Unix(timestamp, 0).Format("2006-01-02 15:04:05"))
}
