package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

func main() {
	// copy()

	copyReadFrom()

	// copyWriteTo()

	copyBuffer()

	//换行输入并打印
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

type Buffer struct {
	bytes.Buffer
	io.ReaderFrom // conflicts with and hides bytes.Buffer's ReaderFrom.
	io.WriterTo   // conflicts with and hides bytes.Buffer's WriterTo.
}

func copy() {
	rb := new(Buffer)
	wb := new(Buffer)

	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		log.Fatalln("Copy did not work properly")
	}

	// log.Println(wb.String())
}

func copyReadFrom() {
	rb := new(Buffer)
	wb := new(bytes.Buffer) // implements ReadFrom.
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		log.Fatalln("Copy did not work properly")
	}
}

func copyWriteTo() {
	rb := new(bytes.Buffer) // implements WriteTo.
	wb := new(Buffer)
	rb.WriteString("hello, world.")
	io.Copy(wb, rb)
	if wb.String() != "hello, world." {
		log.Fatalln("Copy did not work properly")
	}
}

var bufPool sync.Pool

func copyBuffer() {

	buf := bufPool.Get()
	if buf == nil {
		buf = make([]byte, 32*1024)
	}

	rb := new(Buffer)
	wb := new(Buffer)
	rb.WriteString("hello, world.")

	// log.Printf("buf %+v\n", buf.([]byte))
	io.CopyBuffer(wb, rb, buf.([]byte)) // Tiny buffer to keep it honest.
	if wb.String() != "hello, world." {
		log.Fatalln("CopyBuffer did not work properly")
	}

	bufPool.Put(buf)
}
