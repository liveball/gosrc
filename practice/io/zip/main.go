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
)
var (
	urls=[]string{
		"https://ucc.alicdn.com/avatar/img_4220ff7b945d6632ab338ab84f37bfd1.png",
		"https://img.alicdn.com/tfs/TB1t4w5RpXXXXcsXVXXXXXXXXXX-32-32.png",
		"https://img.alicdn.com/tfs/TB1DZMGRpXXXXXHaFXXXXXXXXXX-32-32.png",
	}
)
type zeros struct{}

func (zeros) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

func main() {
	//bigZip := makeZip("big.file", io.LimitReader(zeros{}, 1<<32-1))
	//if err := ioutil.WriteFile("/tmp/big.zip", bigZip, 0666); err != nil {
	//	log.Fatal(err)
	//}

	//biggestZip := makeZip("bigger.zip", bytes.NewReader(biggerZip))
	//if err := ioutil.WriteFile("/tmp/biggest.zip", biggestZip, 0666); err != nil {
	//	log.Fatal(err)
	//}



	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	//var files = []struct {
	//	Name, Body string
	//}{
	//	{"readme.txt", "This archive contains some text files."},
	//	{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	//	{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	//}
	//for _, file := range files {
	//	f, err := w.Create(file.Name)
	//	if err != nil {
	//		log.Fatal(err)it
	//	}
	//	_, err = f.Write([]byte(file.Body))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}

	for _,v:= range urls{
		resp,err:=http.DefaultClient.Get(v)
		if err!=nil{
			log.Fatalln(err)
		}
		defer resp.Body.Close()


		body,err:=ioutil.ReadAll(resp.Body)
		if err!=nil{
			log.Fatalln(err)
		}
		fileName:= url.PathEscape(path.Base(v))
		println(fileName)

		f, err := w.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(body)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}


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


	log.Fatal(http.ListenAndServe(":8080", nil))
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
