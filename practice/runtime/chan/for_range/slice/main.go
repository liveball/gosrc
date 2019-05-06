package main

func main() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	ch := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		ch[i] = make(chan int)

		go func(c []chan int, s []int, j int) {
			println("s[j]:", s[j])
			c[j] <- j
			// close(ch[j])
		}(ch, sli, i)
	}

	for i := range ch {
		println("ch:", <-ch[i])
	}
}
