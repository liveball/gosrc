package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"compress/gzip"

	"github.com/yangjian/lib/log"
)

func main() {
	var err error
	resp, err := http.DefaultClient.Get("http://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll error %s\n", err)
	}
	// fmt.Println(string(body))

	// res, err := http.ReadResponse(b, nil)
	// if err != nil {
	// 	return
	// }

	output := "./"
	// encoding := []string{"gzip"}
	encoding := resp.Header.Get("Content-Encoding")
	base := url.QueryEscape(path.Base("/query?app=macqq"))

	base = path.Join(output, base)
	if len(base) > 250 {
		base = base[:250] + "..."
	}
	if base == output {
		base = path.Join(output, "noname")
	}

	target := base

	n := 0
	for {
		_, err := os.Stat(target)
		//if os.IsNotExist(err) != nil {
		if err != nil {
			log.Warn(err)
			break
		}
		target = fmt.Sprintf("%s-%d", base, n)
		n++
	}

	f, err := os.Create(target)
	if err != nil {
		log.Printf("HTTP-create Cannot create %s: %s\n", target, err)
	}

	var r io.Reader
	r = bytes.NewBuffer(body)

	if len(encoding) > 0 && (encoding == "gzip" || encoding == "deflate") {
		r, err = gzip.NewReader(r)
		if err != nil {
			log.Printf("HTTP-gunzip gzip decode error(%v)", err)
		}
	}

	w, err := io.Copy(f, r)
	if _, ok := r.(*gzip.Reader); ok {
		r.(*gzip.Reader).Close()
	}
	f.Close()

	if err != nil {
		log.Errorf("Saved written(%d) error(%v)", w, err)
	} else {
		log.Printf("Saved written(%d)", w)
	}
}
