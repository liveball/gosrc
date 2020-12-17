package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"runtime"
)

// ResponseWriter： 生成Response的接口

// Handler： 处理请求和生成返回的接口

// ServeMux： 路由，后面会说到ServeMux也是一种Handler

// Conn : 网络连接

const (
	addr = "127.0.0.1:8000"
)

// GODEBUG=gctrace=1 go run main.go |& gcvis
// wrk -c100 -d1m -t4 http://127.0.0.1:8000

type users map[int]string

var visits = expvar.NewInt("visits")

func (u users) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusInternalServerError)
	//return
	//
	//visits.Add(1)
	//
	//fmt.Fprintf(w, "hello world!")
	//
	//for id, name := range u {
	//	fmt.Fprintf(w, "ID(%d),Name(%s)\n", id, name)
	//}
}

func main() {
	fmt.Println("cpu:", runtime.NumCPU(), "goroutine:", runtime.NumGoroutine())
	us := users{
		1: "tom",
		2: "jack",
	}

	err := http.ListenAndServe(addr, us)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
	log.Println("connect ok")

	// http.HandleFunc("/add", testHandler)
	// err := http.ListenAndServe(addr, nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
	// r.ParseForm()
	// params := r.Form
	// fmt.Fprintf(w, params.Get("aid"))
}
