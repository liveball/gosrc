package main

func main() {
	ch := make(chan int)
	ch <- 1

	// ch := make(chan int, 1)

	// var wg sync.WaitGroup
	// for i := 1; i <= 1; i++ {
	// 	wg.Add(1)
	// 	go func(j int) {
	// 		ch <- j
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// close(ch)

	// for i := range ch {
	// 	fmt.Println("ch:", i)
	// }
}

// Act like goroutine called runtime.Gosched.
// casgstatus(gp, _Gwaiting, _Grunning)
// gopreempt_m(gp) // never return

// func goschedImpl(gp *g) {
// 	status := readgstatus(gp)
// 	if status&^_Gscan != _Grunning {
// 		dumpgstatus(gp)
// 		throw("bad g status")
// 	}
// 	casgstatus(gp, _Grunning, _Grunnable)
// 	dropg()
// 	lock(&sched.lock)
// 	globrunqput(gp)
// 	unlock(&sched.lock)

// 	schedule()
// }
