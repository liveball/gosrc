package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//go build -o main main.go && GODEBUG=schedtrace=10000,scheddetail=1 ./main

func main() {
	// godebug()

	t := Teacher{}
	t.ShowA()

	time.Sleep(time.Second * 5)
}

func godebug() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i:", i)
			wg.Done()
		}()
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i:", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}
