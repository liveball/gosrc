package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	consumeLimit = int64(10)
	tokenChan    = make(chan int, consumeLimit) //速率缓冲大小控制
	consumeRate  = int64(1e9 / consumeLimit)

	uchan = make(chan *user) //速率缓冲大小控制
)

func main() {
	// a := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	// go generateToken(time.Duration(consumeRate))
	// for _, v := range a {
	// 	start := time.Now()
	// 	<-tokenChan //消费速率控制
	// 	elapsed := time.Since(start)
	// 	fmt.Println(elapsed)
	// 	fmt.Println(v)
	// }

	// a := []int{11, 12, 13, 14, 15}
	// randomStartTimer(a, testRandomSlice)

	u := &user{}
	setStartTimer(u)
	go func() {
		for u := range uchan {
			// now := time.Now()
			// if int(now.Weekday()) != 4 || now.Hour() != 23 || now.Minute() != 38 {
			// 	time.Sleep(1 * time.Second)
			// 	// println("每周星期4 23点30分运行一次")
			// 	continue
			// }

			// fmt.Println("get now...", now.Format("2006-01-02 15:04:05"))
			// <-getTimer(now, 0, 0, 10).C
			println("get age:", u.age)
		}
	}()

	// time.Sleep(1000 * time.Second)
	select {}
}

func setStartTimer(u *user) {
	go func() {
		for {
			now := time.Now()

			if int(now.Weekday()) != 5 || now.Hour() != 00 || now.Minute() != 21 {
				time.Sleep(1 * time.Second)
				// println("每周星期4 23点15分运行一次")
				continue
			}

			<-getTimer(now, 0, 0, 10).C

			fmt.Println("start now", now.Format("2006-01-02 15:04:05"))

			time.Sleep(20 * time.Second)

			fmt.Println("set now", now.Format("2006-01-02 15:04:05"))

			rand.Seed(now.UnixNano())
			age := rand.Intn(100)
			println("set age:", age)
			u.age = age
			uchan <- u
		}
	}()
}

func getTimer(now time.Time, hour, minute, second int) (t *time.Timer) {
	next := now.Add(time.Hour * time.Duration(hour)).Add(time.Minute * time.Duration(minute)).Add(time.Second * time.Duration(second))
	// next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Hour(), 0, 0, next.Location())
	fmt.Println("next", next.Format("2006-01-02 15:04:05"))
	t = time.NewTimer(next.Sub(now))
	t.Stop()
	return
}

func randomStartTimer(a []int, f func([]int)) {
	go func() {
		for {
			f(a)
			<-getTimer(time.Now(), 0, 0, 5).C
			time.Sleep(1 * time.Second)
		}
	}()
}

func generateToken(duration time.Duration) {
	var (
		timer = time.NewTicker(duration)
		token = 0
	)
	for range timer.C {
		token++
		tokenChan <- token
	}
}

func randomSlice(start int, end int, count int) []int {
	if end < start || (end-start) < count {
		return nil
	}
	nums := make([]int, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		num := r.Intn((end - start)) + start
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func testRandomSlice(a []int) {
	keys := randomSlice(0, len(a), len(a))

	ns := make([]int, 0, len(a))
	for _, k := range keys {
		ns = append(ns, a[k])
	}
}

type user struct {
	age int
}

func (u user) setAge(i int) {
	u.age = i
}
