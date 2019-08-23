package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"unsafe"

	"gosrc/go/src/encoding/gob"
)

type Vector struct {
	X, Y, Z int
}

func main() {
	var network bytes.Buffer // Stand-in for the network.

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)

	msg := `{"name":"依依33399","sex":"男","id_type":"大陆身份证","phone":"156655555","email":"cccc@test.com11","qq":"16635224","pics":{"prove":"http://uat-i0.hdslb.com/bfs/archive/5b1cf41dbd02d1f5f127c3d64eafdd3a22c82afd.jpg_60x60.jpg;http://uat-i0.hdslb.com/bfs/archive/5b1cf41dbd02d1f5f127c3d64eafdd3a22c82afd.jpg_60x60.jpg;http://uat-i0.hdslb.com/bfs/archive/5b1cf41dbd02d1f5f127c3d64eafdd3a22c82afd.jpg_60x60.jpg;http://uat-i0.hdslb.com/bfs/archive/5b1cf41dbd02d1f5f127c3d64eafdd3a22c82afd.jpg_60x60.jpg","positive_side":"http://uat-i0.hdslb.com/bfs/creative/3feb42b6f6782d8aafd01c58a318fa42a8fb4870.png","negative_side":""}}`
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("base64编码:", encoded, "占用空间大小:", unsafe.Sizeof(encoded), "长度:", len(encoded))

	var in bytes.Buffer
	enb := []byte(encoded)
	zl := zlib.NewWriter(&in)
	zl.Write(enb)
	zl.Close()
	fmt.Println("压缩数据:", in.String(), "占用空间大小:", unsafe.Sizeof(in.String()), "长度:", len(in.String()))

	var out bytes.Buffer
	r, _ := zlib.NewReader(&in)
	io.Copy(&out, r)
	fmt.Println("解压数据:", out.String())

	decoded, err := base64.StdEncoding.DecodeString(out.String())
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println("base64解码:", string(decoded))

	zip := DoZlibCompress([]byte("hello, world\n"))
	fmt.Println(zip)
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	fmt.Println(string(DoZlibUnCompress(buff)))
}

//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, err := zlib.NewReader(b)
	if err!=nil{
		fmt.Println(err)
	}
	io.Copy(&out, r)
	r.Close()
	return out.Bytes()
}