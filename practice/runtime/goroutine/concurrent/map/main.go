package main

var (
	m = make(map[int]int)
)

func main() {
	defer func() { <-make(chan bool) }()
	go func() {
		for {
			m[1] = 1
		}
	}()
	go func() {
		for {
			_ = m[1]
		}
	}()
}
