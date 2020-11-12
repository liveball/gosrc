package stages

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
)

func serviceConnLogSNI(conn net.Conn) {
	defer conn.Close()

	log.Printf("receive a conn from addr(%v)", conn.RemoteAddr())

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))

	var buf bytes.Buffer
	if _, err := io.CopyN(&buf, conn, 1+2+2); err != nil {
		log.Printf("failed to read record header error(%v)\n", err)
	}

	length := binary.BigEndian.Uint16(buf.Bytes()[3:5])
	if _, err := io.CopyN(&buf, conn, int64(length)); err != nil {
		log.Printf("failed to read client hello record error(%v)\n", err)
		return
	}

	ch, ok := ParseClientHello(buf.Bytes())
	if !ok {
		log.Printf("falied to parse client hello\n")
	} else {
		log.Printf("received connection to sni(%q)!", ch.SNI)
	}
}
