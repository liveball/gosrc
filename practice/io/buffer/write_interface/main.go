package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	var i interface{}

	var buf bytes.Buffer
	i = 65
	buf.WriteString(string(i.(int)))
	fmt.Println("buf:", buf.Bytes(), string(buf.Bytes()), buf.String()) // [1]

	var buf2 bytes.Buffer
	buf2.WriteString(strconv.Itoa(i.(int)))
	fmt.Println("buf2:", buf2.Bytes(), string(buf2.Bytes()), buf2.String()) // [49]

	var res int32
	a := make([]byte, 4)
	a[2] = buf2.Bytes()[0]
	a[3] = buf2.Bytes()[1]
	err := binary.Read(bytes.NewBuffer(a), binary.BigEndian, &res)
	if err != nil {
		fmt.Printf("binary.Read data(%v) error(%+v)", buf2.Bytes(), err)
	}
	fmt.Println(a, int(res))
}

//字符1 accii表对应10进制是49
