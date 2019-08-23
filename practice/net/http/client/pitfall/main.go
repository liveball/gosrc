package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 10)
	}))
	defer svr.Close()

	fmt.Println("making request")
	start := time.Now()
	http.Get(svr.URL)

	fmt.Println("finished request", time.Now().Sub(start))
}
