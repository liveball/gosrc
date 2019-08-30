package main

func main() { // breakpoint 1
	ch := make(chan int)
	go func() {
		for i := range ch {
			println(i) // breakpoint 2
		}
	}()

	ch <- 1

	wait := make(chan int) // breakpoint 3
	<-wait
}