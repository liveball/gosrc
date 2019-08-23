package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

var (
	client *http.Client

	dialTimeout = time.Duration(100 * time.Millisecond)
	rwTimeout   = time.Duration(300 * time.Millisecond)

	cxtTimeout = time.Duration(200 * time.Millisecond)

	link = "https://avatar-static.segmentfault.com/339/518/3395183340-2968_big64"
)

func main() {
	//client= NewDialTimeoutClient()//拨号超时
	client = NewConnRWTimeoutClient(dialTimeout, rwTimeout) //连接 读写超时

	ctx, cancel := context.WithTimeout(context.Background(), cxtTimeout)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		log.Printf("http.NewRequest error(%v)", err)
		return
	}

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Printf("client.Do error(%v)", err)
		return
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll error(%v)", err)
		return
	}

	log.Println(len(bs))
}

func dialTimeoutFunc(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, dialTimeout)
}

func NewDialTimeoutClient() *http.Client {
	transport := http.Transport{
		Dial: dialTimeoutFunc,
	}

	return &http.Client{
		Transport: &transport,
	}
}

func timeOutDialer(dialTimeout time.Duration, rwTimeout time.Duration) func(network, addr string) (c net.Conn, err error) {
	return func(network, addr string) (c net.Conn, err error) {
		conn, err := net.DialTimeout(network, addr, dialTimeout)
		if err != nil {
			return nil, err
		}

		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func NewConnRWTimeoutClient(dialTimeout time.Duration, rwTimeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: timeOutDialer(dialTimeout, rwTimeout),
		},
	}
}
