package main

func main() {
	ch := make(chan int)
	ch <- 1

	// var wg sync.WaitGroup
	// for i := 1; i < 1; i++ {
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
