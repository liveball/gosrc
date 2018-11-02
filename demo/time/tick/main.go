package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/davecgh/go-spew/spew"
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

	a := []int{11, 12, 13, 14, 15}
	randomStartTimer(a, testRandomSlice)

	// u := &user{}
	// setStartTimer(u)
	// go func() {
	// 	for u := range uchan {
	// 		println("get age:", u.age)
	// 	}
	// }()

	time.Sleep(10 * time.Second)
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
	// spew.Dump(keys)

	ns := make([]int, 0, len(a))
	for _, k := range keys {
		ns = append(ns, a[k])
	}
	spew.Dump(ns)
}

func randomStartTimer(a []int, f func([]int)) {
	go func() {
		for {
			f(a)
			<-getTimer().C
		}
	}()
}

func setStartTimer(u *user) {
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			age := rand.Intn(100)
			println("set age:", age)

			u.age = age
			uchan <- u
			<-getTimer().C
		}
	}()
}

func getTimer() (t *time.Timer) {
	now := time.Now()
	next := now.Add(time.Hour * 24)

	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

	fmt.Println("next", next.Format("2006-01-02 15:04:05"))
	t = time.NewTimer(next.Sub(now))
	// t = time.NewTimer(1 * time.Second)
	return
}

type user struct {
	age int
}

func (u user) setAge(i int) {
	u.age = i
}
