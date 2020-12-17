package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"runtime/trace"
	"time"
)

func main() {
	//创建输出文件
	f0, err := os.Create("trace.out")
	defer f0.Close()
	if err != nil {
		log.Fatalf("could not create trace.out:", err)
	}

	//开始采集memory的信息
	if err := trace.Start(f0); err != nil {
		log.Fatalf("could not start trace :", err)
	}
	defer trace.Stop()

	//创建输出文件
	f, err := os.Create("cpu.prof")
	defer f.Close()
	if err != nil {
		log.Fatalf("could not create cpu profile:", err)
	}

	//开始采集cpu的信息
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("could not start cpu profile:", err)
	}
	defer pprof.StopCPUProfile()

	//test code ----------start
	testCall() //test
	testGoroutine()
	//test code ----------end

	//runtime.GC()//回收内存  762.95MB =》1.16MB

	collectMemory()
	collectGoroutine()
}

func testGoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	for n := 0; n < 5; n++ {
		i := n

		go func() {
			log.Println(i)

			select {
			case <-ctx.Done():
				log.Println("ctx.Done()")
				break
			default:
				log.Println("default")
			}

		}()
	}

	cancel()
}

func collectMemory() {
	//创建输出文件
	f, err := os.Create("mem.prof")
	defer f.Close()
	if err != nil {
		log.Fatalf("could not create mem profile:", err)
	}

	//开始采集memory的信息
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatalf("could not start mem profile:", err)
	}
}

func collectGoroutine() {
	//创建输出文件
	f, err := os.Create("goroutine.prof")
	defer f.Close()
	if err != nil {
		log.Fatalf("could not create goroutine profile:", err)
	}

	//开始采集memory的信息
	if pf := pprof.Lookup("goroutine"); pf == nil {
		log.Fatalf("could not start goroutine profile:", err)
	} else {
		pf.WriteTo(f, 0)
	}
}

func testCall() {
	// 主逻辑区，进行一些简单的代码运算
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

}

const (
	col = 10000
	row = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(100000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}
