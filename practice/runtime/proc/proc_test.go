package main

import (
	"runtime"
	"testing"
	"time"
)

func TestStopTheWorldDeadlock(t *testing.T) {
	StopTheWorldDeadlock()
}

func TestYieldProgress(t *testing.T) {
	testYieldProgress(false)
}

func TestYieldLockedProgress(t *testing.T) {
	testYieldProgress(true)
}


func TestYieldLocked(t *testing.T) {
	const N = 10
	c := make(chan bool)
	go func() {
		runtime.LockOSThread()
		for i := 0; i < N; i++ {
			runtime.Gosched()
			time.Sleep(time.Millisecond)
		}
		c <- true
		// runtime.UnlockOSThread() is deliberately omitted
	}()
	<-c
}
