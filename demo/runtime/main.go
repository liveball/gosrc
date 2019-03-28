package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gosrc/demo/runtime/sched"
)

// go tool compile -N -l -S main.go > main.s

//GODEBUG=schedtrace=10000,scheddetail=1 ./main
func main() {
	// println(1)

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	println(2)
	// 	wg.Done()
	// }()

	// wg.Wait()
	// Set this threshold low for demonstration purposes.
	sched.OversleepThreshold = time.Microsecond
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

type warn struct{}

func (w *warn) Warningf(string, ...interface{}) {
	fmt.Println("111")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// ctx := appengine.NewContext(r)
	sched.Check(&warn{})
}
