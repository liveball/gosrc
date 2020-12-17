package stages

import (
	"io"
	"log"
	"net"
)

func serviceProxy(conn net.Conn) {
	defer conn.Close()

	log.Printf("receive a conn from addr(%v)", conn.RemoteAddr())

	upstream, err := net.Dial("tcp", "gophercon.com:https")//gophercon.com:https
	if err != nil {
		log.Println(err)
		return
	}
	defer upstream.Close()

	go io.Copy(upstream, conn)

	n, err := io.Copy(conn, upstream)
	log.Printf("conn fin with n(%d) error(%v)", n, err)
}
