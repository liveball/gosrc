package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"
)

var (
	urls = []string{
		"//www.jianshu.com/p/b63095c59361",
		"https://tech.meituan.com/2017/05/19/about-desk-io.html?token=11",
		"https://tech.meituan.com/2017/05/19/about-desk-io.html?token=dssd",
	}
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	pro := "https:"

	wg.Add(len(urls))
	for _, v := range urls {

		if _, err := url.ParseRequestURI(v); err != nil {
			log.Printf("url.ParseRequestURI url(%s) error(%v)", v, err)
			continue
		}

		var vv string
		if !strings.Contains(v, pro) {
			vv = pro + v
		} else {
			vv = v
		}

		println(url.QueryEscape(path.Base(vv)))

		if strings.Contains(v, "?") {
			println("del token", v[:strings.Index(v, "?")])
		}

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
			defer func() {
				cancel()
				wg.Done()
			}()

			req, err := http.NewRequest("GET", vv, nil)
			if err != nil {
				log.Printf("http.NewRequest (%v)\n", err)
				return
			}
			req = req.WithContext(ctx)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("http.DefaultClient.Do (%v)\n", err)
				return
			}
			defer resp.Body.Close()

			log.Println(req.URL.Path, resp.StatusCode)
		}()
	}
	wg.Wait()

	log.Println(time.Now().Sub(start))
}
