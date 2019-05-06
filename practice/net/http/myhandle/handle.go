package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type users map[int]string

//显示实现Handler接口
func (u users) ServeHTTP(w ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	for id, name := range u {
		fmt.Fprintf(w, "ID(%d),Name(%s)\n", id, name)
	}
}

func (u users) list(w ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	for id, name := range u {
		fmt.Fprintf(w, "ID(%d),Name(%s)\n", id, name)
	}
}

func main() {
	user := users{
		1: "tom",
		2: "jack",
	}

	r := new(http.Request)
	r.URL = &url.URL{
		Path: "/list",
	}
	// fmt.Printf("Request(%+v)\n", r)

	w := &response{
		req: r,
	}

	//1、用户自定义实现
	// user.ServeHTTP(w, r)

	//2、使用  DefaultServeMux HandleFunc 注册
	// HandleFunc("/list", user.list)

	//3、使用  DefaultServeMux Handle 注册
	// Handle("/list", HandlerFunc(user.list))

	//4、使用new ServeMux 注册
	mux := new(ServeMux)
	mux.HandleFunc("/list", HandlerFunc(user.list))
	// mux.Handle("/list", HandlerFunc(user.list))
	serverHandler{&Server{Addr: "127.0.0.1", Handler: mux}}.ServeHTTP(w, r)
}

type response struct {
	req    *http.Request
	status int // status code passed to WriteHeader
}

func (r *response) Write(data []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	w := bufio.NewWriter(buf)
	n, err = w.Write(data)
	fmt.Println(string(data), n)
	return
}

func (w *response) WriteHeader(code int) {
	w.status = code
	fmt.Println(w.status)
}

//ResponseWriter for test http handle
type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
}

//DefaultServeMux for default
var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

type serverHandler struct {
	srv *Server
}

type Server struct {
	Addr    string
	Handler Handler
}

func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *http.Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	fmt.Printf("exec ServeHTTP(%+v)\n", handler)
	handler.ServeHTTP(rw, req)
}

//Handler for fac interface
type Handler interface {
	ServeHTTP(ResponseWriter, *http.Request)
}

//HandlerFunc  隐式实现Handler接口
type HandlerFunc func(ResponseWriter, *http.Request)

//ServeHTTP implement Handler
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *http.Request) {
	fmt.Printf("implement Handle ResponseWriter(%+v) Request(%+v\n", w, r)
	f(w, r)
}

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

// HandleFunc registers the handler function for the given pattern  in the DefaultServeMux.
func HandleFunc(pattern string, handler func(ResponseWriter, *http.Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

//ServeMux by user define
type ServeMux struct {
	m map[string]muxEntry
}

type muxEntry struct {
	h       Handler
	pattern string
}

//Handle for user ServeMux
func (mux *ServeMux) Handle(pattern string, Handler Handler) {
	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	mux.m[pattern] = muxEntry{h: Handler, pattern: pattern}
	fmt.Printf("register Handle mux.m (%+v)\n", mux.m)
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *http.Request)) {
	mux.Handle(pattern, HandlerFunc(handler))
}

//ServeHTTP for user ServeMux ServeHTTP
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *http.Request) {
	h, _ := mux.Handler(r)
	fmt.Printf("get ServeHTTP Handler(%+v)\n", h)
	h.ServeHTTP(w, r)
}

// stripHostPort returns h without any trailing ":<port>".
func stripHostPort(h string) string {
	// If no port on host, return unchanged
	if strings.IndexByte(h, ':') == -1 {
		return h
	}
	host, _, err := net.SplitHostPort(h)
	if err != nil {
		return h // on error, return unchanged
	}
	return host
}

//Handler for user ServeMux handler
func (mux *ServeMux) Handler(r *http.Request) (h Handler, pattern string) {
	host := stripHostPort(r.Host)
	return mux.handler(host, r.URL.Path)
}

func (mux *ServeMux) handler(host, path string) (h Handler, pattern string) {
	if h == nil {
		h, pattern = mux.match(path)
		fmt.Printf("mux.match handler(%+v)\n", h)
	}
	return
}

func (mux *ServeMux) match(path string) (h Handler, pattern string) {
	// Check for exact match first.
	v, ok := mux.m[path]
	if ok {
		fmt.Printf("match v(%+v) ok(%+v)\n", v, ok)
		return v.h, v.pattern
	}
	// Check for longest valid match.
	var n = 0
	for k, v := range mux.m {
		if !pathMatch(k, path) {
			continue
		}
		if h == nil || len(k) > n {
			n = len(k)
			h = v.h
			pattern = v.pattern
		}
	}
	return
}

func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		return false
	}
	n := len(pattern)
	if pattern[n-1] != '/' {
		return pattern == path
	}
	return len(path) >= n && path[0:n] == pattern
}
