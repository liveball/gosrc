package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"
)

//从网络上获取多张图片并打包下载
var (
	urls=[]string{
		"https://ucc.alicdn.com/avatar/img_4220ff7b945d6632ab338ab84f37bfd1.png",
		"https://img.alicdn.com/tfs/TB1t4w5RpXXXXcsXVXXXXXXXXXX-32-32.png",
		"https://img.alicdn.com/tfs/TB1DZMGRpXXXXXHaFXXXXXXXXXX-32-32.png",
	}
)

var (
	filename="test"
)

//直接下载附件 http://localhost:8080/dd

//文件列表  http://localhost:8080/files

//返回下载链接  http://localhost:8080/download

func main() {

	//直接下载可能包含的附件 direct download => dd
	http.HandleFunc("/dd", func(w http.ResponseWriter, r *http.Request) {
		var err error
		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", filename))

		zipW := zip.NewWriter(w)
		defer zipW.Close()
		if err = attachmentDownload(zipW, urls); err != nil {
			fmt.Fprint(w,err)
			return
		}
	})


	fileHandle()//文件服务

	log.Println("server start:")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//attachmentDownload download urls
func attachmentDownload(w *zip.Writer, urls []string) (err error){
	var (
		wg    sync.WaitGroup
		mux sync.Mutex
	)

	wg.Add(len(urls))
	for _, v := range urls {
		vv := v
		go func() {
			mux.Lock()
			packCompress(w, vv)
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	if err = w.Close();err != nil {
		log.Fatalf("AttachmentDownload w.Close() urls(%+v) error(%v)", urls, err)
		return
	}
	return
}


func packCompress(w *zip.Writer, addr string) {
	resp, err := http.DefaultClient.Get(addr)
	if err != nil {
		log.Fatalf("packCompress http.DefaultClient.Get addr(%+v) error(%v)", addr, err)
		return
	}
	defer resp.Body.Close()

	base := url.QueryEscape(path.Base(addr))
	if len(base) > 50 {
		base = base[:50] + "..."
	}

	fileName := base
	n := 0
	for {
		_, err := os.Stat(fileName)
		if err != nil {
			break
		}
		fileName = fmt.Sprintf("%s-%d", base, n)
		n++
	}

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatalf("packCompress ioutil.ReadAll addr(%+v) error(%v)", addr, err)
		return
	}

	f, err := w.Create(fileName)
	if err != nil {
		log.Fatalf("packCompress w.Create addr(%+v) error(%v)", addr, err)
		return
	}

	_, err = f.Write(body)
	if err != nil {
		log.Fatalf("packCompress f.Write addr(%+v) error(%v)", addr, err)
		return
	}
}



//1、服务器保存zip文件
//2、可查询文件列表
//3、返回下载链家
func fileHandle(){
	var (
		wg    sync.WaitGroup
		mux sync.Mutex
	)

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	wg.Add(len(urls))
	for _, v := range urls {
		vv := v
		go func() {
			mux.Lock()
			packCompress(w, vv)
			mux.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()


	zipFile:="file.zip"
	f, err := os.OpenFile(zipFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("save os.OpenFile urls(%+v) error(%v)", urls, err)
		return
	}
	if _,err=buf.WriteTo(f);err != nil {
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

		host:=strings.Join([]string{scheme, r.Host, "/files"}, "")

		dw:="<a href=" + host+"/"+zipFile+">下载</a>"
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
