package stages

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func AcceptLoop() {
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err, ok := err.(net.Error); ok && err.Temporary() {
			log.Printf("temporary accept error:%v; sleeping 1s...")
			time.Sleep(time.Second * 1)
			continue
		} else if err != nil {
			log.Fatal(err)
		}

		//go serviceConn(conn)
		//go copyToStderr(conn)
		//go serviceProxy(conn)
		//go serviceConnLogSNI(conn)
		//go serviceConnProxyAndSNI(conn)
		go serviceConnTLS(conn)
	}
}

func serviceConn(conn net.Conn) {
	defer conn.Close()

	n, err := io.Copy(os.Stderr, conn)
	log.Printf("copied n(%d) bytes and ended with error(%v)", n, err)
}
