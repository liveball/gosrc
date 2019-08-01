package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"net/http"
	_ "net/http/pprof"
)

type p struct {
	i int
}

var (
	jobNum  = 100
)

func main() {
	for i := 0; i < jobNum; i++ {
		go func(i int) {
			dispatch()
		}(i)
	}

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func dispatch() {
	for {
		res := foo()
		fmt.Printf("%p\n", res)
		tks := make([]*p, 0, len(res))
		for _, v := range res {
			tks = append(tks, v)
		}

		time.Sleep(time.Millisecond * 20)
	}
}

func foo() (res []*p) {
	res = make([]*p, 0, 1)
	res = append(res, &p{i: 100})
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")

	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
)

func init() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}
