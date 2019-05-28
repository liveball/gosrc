package main

import (
	"log"
	"os"
)

func main() {
	for {
		//不关闭fd
		_, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalln(err) //open a.txt: too many open files in system
		}

		//关闭fd
		// f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE, 0600)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// f.Close()
	}
}
