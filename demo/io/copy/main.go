package main

import (
	"bytes"
	"container/list"
	"encoding/gob"
	"fmt"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/mohae/deepcopy"
)

//Conn for queue
type Conn struct {
	queue list.List
}

func (c *Conn) Write(p []byte) (n int, err error) {
	c.queue.PushBack(p) //定时器发送到网络对端
	return len(p), nil
}

func main() {
	// var err error
	// resp, err := http.DefaultClient.Get("http://www.baidu.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()
	f, err := os.Open("a.txt")
	// f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// io.Copy(f, resp.Body) //写文件

	conn := &Conn{}
	// io.Copy(conn, f)
	// deepCopy(conn, f)

	conn = deepcopy.Copy(f).(*Conn)
	if !reflect.DeepEqual(f, conn) {
		fmt.Printf("got %#v; want %#v\n", conn, f)
		return
	}

	spew.Dump(conn)
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
