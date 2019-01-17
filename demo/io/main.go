package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.OpenFile("./a.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf *bufio.Reader
	buf = bufio.NewReader(f)
	for {
		str, err := buf.ReadString('\n') //log agent
		if err == io.EOF {
			fmt.Printf("exit error(%v)\n", err)
			return
		}

		fmt.Println(str)
		data := strings.Split(str, ",")
		if len(data) == 0 {
			return
		}

		// fmt.Println(data)
		if data[0] != "" && data[1] != "" {
			fmt.Printf("data[0](%s) data[1](%s) time(%s)\n", data[0], data[1], time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}
