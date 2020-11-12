package stages

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
)

func serviceConnTLS(conn net.Conn){
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

	//TODO move to main()

	cert,err:=tls.LoadX509KeyPair("localhost.pem", "localhost-key.pem")
	if err!=nil{
		log.Fatal(err)
	}
	config:=&tls.Config{
		Certificates:[]tls.Certificate{cert},
	}
	c:=tls.Server(prefixConn{
		Reader: io.MultiReader(&buf, conn),
		Conn:   conn,
	}, config)


     //copyToStderr(c)

     proxyConn(c, "gophercon.com:80")
}
