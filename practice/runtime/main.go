package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"gosrc/practice/runtime/sched"
)

// go tool compile -N -l -S main.go > main.s

//GODEBUG=schedtrace=10000,scheddetail=1 ./main

// curl http://127.0.0.1:8000/debug/pprof/trace?seconds=20 > trace.out
// go tool trace trace.out

func main() {
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
