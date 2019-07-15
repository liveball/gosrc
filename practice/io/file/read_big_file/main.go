package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/data/app/go/src/gosrc/practice/io/file/read_big_file/a.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	for scanner.Scan() {
		log.Println(scanner.Text())
	}
}
