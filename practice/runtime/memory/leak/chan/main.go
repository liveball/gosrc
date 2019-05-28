package main

func main() {
	for {

		//不关闭chan
		_ = make(chan int)

		//关闭chan
		// ch := make(chan int)
		// close(ch)

	}
}
