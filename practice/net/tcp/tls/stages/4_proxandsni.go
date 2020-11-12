package stages

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
)

func serviceConnProxyAndSNI(conn net.Conn) {
	defer conn.Close()

	log.Printf("receive a conn from addr(%v)", conn.RemoteAddr())

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, conn, 1+2+2); err != nil {
		log.Printf("failed to read record header error(%v)\n", err)
		conn.Close()
		return
	}

	length := binary.BigEndian.Uint16(buf.Bytes()[3:5])
	if _, err := io.CopyN(&buf, conn, int64(length)); err != nil {
		log.Printf("failed to read client hello record error(%v)\n", err)
		conn.Close()
		return
	}

	ch, ok := ParseClientHello(buf.Bytes())
	if !ok {
		log.Printf("failed to parse client hello\n")
	} else {
		log.Printf("received connection for sni(%q)\n", ch.SNI)
	}

	conn.SetDeadline(time.Time{}) //reset deadline
	conn.(*net.TCPConn).SetKeepAlive(true)
	conn.(*net.TCPConn).SetKeepAlivePeriod(3 * time.Minute)

	proxyConn(
		prefixConn{
			Reader: io.MultiReader(&buf, conn),
			Conn:   conn,
		}, "gophercon.com:https")
}

type prefixConn struct {
	io.Reader
	net.Conn
}

func (c prefixConn) Read(b []byte) (int, error) {
	return c.Reader.Read(b)
}

func proxyConn(conn net.Conn, addr string) {
	upstream, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("net.Dial error(%v)\n", err)
		return
	}

	defer upstream.Close()

	go io.Copy(upstream, conn)

	n, err := io.Copy(conn, upstream) //splice from 1.11
	log.Printf("proxy connection finished with n(%v)bytes error(%v)\n", n, err)
}
