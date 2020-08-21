package main

import (
	"testing"
)

func TestNewPool_Acquire_Release(t *testing.T) {
	p := NewPool(5)
	c,err:= p.Acquire()
	if err!=nil{
		t.Error(err)
		return
	}

    n,err:= c.Write([]byte("hello"))
    if err!=nil{
    	t.Error(err)
		return
	}

    t.Log(n)
}
