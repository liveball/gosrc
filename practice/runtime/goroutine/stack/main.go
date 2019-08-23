package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

//awk -F '[' '/goroutine \d*/{print "[" $2}' stack.log | sort | uniq -c | sort -k1nr | head -20

//2 [running]:
//2 [sleep]:
//2 [syscall]:

func main() {
	//SetupStackTrap()
	SetupStackTrap("stack.log")
	time.Sleep(100 * time.Second)
}

// how to use?
// kill -USR1 pid
// tail stack.log

const (
	timeFormat = "2006-01-02 15:04:05"
)

var (
	stdFile = "./stack.log"
)

func SetupStackTrap(args ...string) {
	if len(args) > 0 {
		stdFile = args[0]
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			dumpStacks()
		}
	}()
}

func dumpStacks() {
	buf := make([]byte, 1638400)
	buf = buf[:runtime.Stack(buf, true)]
	writeStack(buf)
}

func writeStack(buf []byte) {
	fd, _ := os.OpenFile(stdFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	now := time.Now().Format(timeFormat)
	fd.WriteString("\n\n\n\n\n")
	fd.WriteString(now + " stdout:" + "\n\n")
	fd.Write(buf)
	fd.Close()
}
