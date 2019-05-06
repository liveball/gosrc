package main

import "time"

func myprint(i int) {
	// fmt.Println("hello", i)
	println("hello", i)
}

func main() {
	c := make(chan struct{}, 1)
	c <- struct{}{}
	for i := 0; i < 20; i++ {
		go func() {
			defer func() {
				c <- struct{}{}
			}()
			println(&i)
			myprint(i)
		}()
		time.Sleep(time.Millisecond)

		// go func(j int) {
		// 	defer func() {
		// 		c <- struct{}{}
		// 	}()
		// 	println(&j)
		// 	myprint(j)
		// }(i)

		// j := i
		// go func() {
		// 	defer func() {
		// 		c <- struct{}{}
		// 	}()
		// 	// myprint(i)
		// 	println(&j)
		// 	myprint(j)
		// }()

	}

	for i := 0; i < 20; i++ {
		<-c
	}
}
