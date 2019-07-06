package main


func main() {
	defer func() { <-make(chan bool) }()

	m := make(map[int]int)
	go func() {
			m[1] = 1
	}()

	for _, v := range m {
		println(v)
	}
}
