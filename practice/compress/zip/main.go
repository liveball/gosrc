package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

//从网络上获取多张图片并打包下载
var (
	urls = []string{
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
		"https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64",
	}
)

var (
	filename = "test"
	client   *http.Client

	Timeout         = 3000
	Dial            = 1
	KeepAlive       = 60
	IdleConnTimeout = 90

	MaxIdleConns        = 10
	MaxIdleConnsPerHost = 10

	ReqContextTimeout = 1 //下载上下文超时
)

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(Dial) * time.Second,
				KeepAlive: time.Duration(KeepAlive) * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
		},
		Timeout: time.Duration(Timeout) * time.Millisecond,
	}

	log.Println("设置超时时间：", client.Timeout)
}

//直接下载附件 http://localhost:8080/dl

//文件列表  http://localhost:8080/files

//返回下载链接  http://localhost:8080/download

func main() {

	dl() //直接下载

	//fileHandle() //文件服务
	log.Println("server start")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dl() {
	//直接下载可能包含的附件 direct download => dd
	http.HandleFunc("/dl", func(w http.ResponseWriter, r *http.Request) {
		var err error
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", filename))

		zipW := zip.NewWriter(w)
		defer func() {
			if err = zipW.Close(); err != nil {
				fmt.Fprint(w, err)
			}
		}()

		//if err = attachmentDownload(zipW, urls); err != nil {
		//	fmt.Fprint(w, err)
		//	return
		//}

		start := time.Now()

		if err = concurrentAttachmentDownload(zipW, urls); err != nil { //并发
			fmt.Fprint(w, err)
			return
		}

		log.Println("耗时：", time.Now().Sub(start))

	})
}

//concurrentAttachmentDownload download urls
func concurrentAttachmentDownload(w *zip.Writer, urls []string) (err error) {
	var (
		wg  sync.WaitGroup
		mux sync.Mutex
	)

	wg.Add(len(urls))

	checkNameMap := make(map[string]struct{})
	for i, v := range urls {
		vv := v
		n := i
		go func() {
			defer func() {
				wg.Done()
				//log.Printf("goroutine exit at fetch url(%s)", vv)
			}()
			mux.Lock()
			packCompress(w, checkNameMap, n, vv)
			mux.Unlock()

		}()
	}
	wg.Wait()
	return
}

//attachmentDownload download urls
func attachmentDownload(w *zip.Writer, urls []string) (err error) {
	checkNameMap := make(map[string]struct{})
	for i, v := range urls {
		n := i
		packCompress(w, checkNameMap, n, v)
	}
	return
}

func packCompress(w *zip.Writer, checkNameMap map[string]struct{}, n int, addr string) {
	req, err := http.NewRequest("GET", addr, nil)
	if err != nil {
		log.Printf("http.NewRequest (%v)\n", err)
		return
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ReqContextTimeout)*time.Second)
	//defer cancel()
	//req = req.WithContext(ctx)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do (%v)\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("packCompress ioutil.ReadAll addr(%+v) error(%v)", addr, err)
		return
	}

	fileType := http.DetectContentType(body)
	if fileType != "image/jpg" && fileType != "image/jpeg" && fileType != "image/png" {
		log.Printf("packCompress http.DetectContentType filetype(%s) addr(%+v) error(%v)", fileType, addr, err)
		return
	}

	naddr := subStr(addr, "?")
	base := url.QueryEscape(path.Base(naddr))
	fileName := base

	if _, ok := checkNameMap[fileName]; ok {
		sls := strings.Split(base, ".")
		if len(sls) == 2 {
			fn := fmt.Sprintf("%s-%d", sls[0], n)
			fileName = fn + "." + sls[1]
		}
	}
	checkNameMap[fileName] = struct{}{}

	f, err := w.Create(fileName)
	if err != nil {
		log.Printf("packCompress w.Create addr(%+v) error(%v)", addr, err)
		return
	}

	_, err = f.Write(body)
	if err != nil {
		log.Printf("packCompress f.Write addr(%+v) error(%v)", addr, err)
		return
	}
}

func subStr(all, sub string) (res string) {
	if all == "" || sub == "" {
		return
	}
	if strings.Contains(all, sub) {
		res = all[:strings.Index(all, sub)]
	} else {
		res = all
	}
	return
}

//1、服务器保存zip文件
//2、可查询文件列表
//3、返回下载链家
func fileHandle() {
	var (
		wg  sync.WaitGroup
		mux sync.Mutex
	)

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	wg.Add(len(urls))

	checkNameMap := make(map[string]struct{})
	for i, v := range urls {
		vv := v
		n := i
		go func() {
			mux.Lock()
			packCompress(w, checkNameMap, n, vv)
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	zipFile := "file.zip"
	f, err := os.OpenFile(zipFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("save os.OpenFile urls(%+v) error(%v)", urls, err)
		return
	}
	if _, err = buf.WriteTo(f); err != nil {
		log.Fatalf("save os.OpenFile urls(%+v) error(%v)", urls, err)
	}

	uploadPath := "./"
	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Proto)
		// output:HTTP/1.1
		fmt.Println(r.TLS)
		// output: <nil>
		fmt.Println(r.Host)
		// output: localhost:9090
		fmt.Println(r.RequestURI)
		// output: /index?id=1

		scheme := "http://"
		if r.TLS != nil {
			scheme = "https://"
		}

		host := strings.Join([]string{scheme, r.Host, "/files"}, "")

		dw := "<a href=" + host + "/" + zipFile + ">下载</a>"
		fmt.Println(dw)
		fmt.Fprintln(w, dw)
	})
}

func makeZip(name string, r io.Reader) []byte {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	wf, err := w.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = io.Copy(wf, r); err != nil {
		log.Fatal(err)
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}
