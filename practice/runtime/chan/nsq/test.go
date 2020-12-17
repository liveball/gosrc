package main

import "fmt"

func main()  {
	messagePumpStartedChan := make(chan bool)
	go messagePump("start", messagePumpStartedChan)
	<-messagePumpStartedChan
}

func messagePump(msg string, startedChan chan bool)  {
    fmt.Println(msg)
    close(startedChan)
}