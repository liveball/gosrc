package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

const (
	addr1 = "localhost:8060"
	addr2 = "localhost:8061"
)

type customReader struct {
	iter int
	size int
}

func (r *customReader) Read(b []byte) (int, error) {
	maxRead := r.size - r.iter
	if maxRead == 0 {
		return 0, io.EOF
	}
	if maxRead > len(b) {
		maxRead = len(b)
	}
	n, err := rand.Read(b[:maxRead])
	r.iter += maxRead
	return n, err
}

func (r *customReader) Close() error {
	// Uncomment this for even more races :((
	// r.iter = 0
	return nil
}

func (r *customReader) Reset() {
	r.iter = 0
}

func second(w http.ResponseWriter, r *http.Request) {
	ioutil.ReadAll(r.Body)
}

func first(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://"+addr2+"/second", http.StatusTemporaryRedirect)
}

func main() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/first", first)
	srv1 := http.Server{
		Addr:    addr1,
		Handler: mux1,
	}
	go func() {
		err := srv1.ListenAndServe()
		fmt.Printf("err: %v", err)
	}()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/second", second)
	srv2 := http.Server{
		Addr:    addr2,
		Handler: mux2,
	}
	go func() {
		err := srv2.ListenAndServe()
		fmt.Printf("err: %v", err)
	}()

	time.Sleep(time.Second)
	client := http.DefaultClient

	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		reader := &customReader{size: 900000}
		wg.Add(1)
		go func(r *customReader) {
			for {
				req, err := http.NewRequest(http.MethodPut, "http://"+addr1+"/first", r)
				if err != nil {
					fmt.Printf("%v", err)
				}
				req.GetBody = func() (io.ReadCloser, error) {
					return &customReader{size: 900000}, nil
				}

				if resp, err := client.Do(req); err != nil {
					// fmt.Printf("error: %v", err)
				} else if resp.StatusCode >= http.StatusBadRequest {
					// fmt.Printf("status code: %d", resp.StatusCode)
				}

				// Reset reader and try to reuse it in next request
				r.Reset()
			}

			wg.Done()
		}(reader)
	}

	wg.Wait() // infinite wait

	srv1.Close()
	srv2.Close()
}
