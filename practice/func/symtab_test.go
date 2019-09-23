package main

import (
	"gosrc/go/src/fmt"
	"runtime"
	"strings"
	"testing"
)

func TestCaller(t *testing.T) {
	procs := runtime.GOMAXPROCS(-1)

	fmt.Println(procs)

	procs=1
	c := make(chan bool, procs)
	for p := 0; p < procs; p++ {
		go func() {
			for i := 0; i < 2; i++ {
				testCallerFoo(t)
			}
			c <- true
		}()
		defer func() {
			<-c
		}()
	}
}

// These are marked noinline so that we can use FuncForPC
// in testCallerBar.
//go:noinline
func testCallerFoo(t *testing.T) {
	testCallerBar(t)
}

//go:noinline
func testCallerBar(t *testing.T) {
	for i := 0; i < 2; i++ {
		pc, file, line, ok := runtime.Caller(i)
		f := runtime.FuncForPC(pc)

		t.Logf("symbol info %d: %t %d %d %s %s %d",
			i, ok, f.Entry(), pc, f.Name(), file, line)


		if !ok ||
			!strings.HasSuffix(file, "symtab_test.go") ||
			(i == 0 && !strings.HasSuffix(f.Name(), "testCallerBar")) ||
			(i == 1 && !strings.HasSuffix(f.Name(), "testCallerFoo")) ||
			line < 5 || line > 1000 ||
			f.Entry() >= pc {
			t.Errorf("incorrect symbol info %d: %t %d %d %s %s %d",
				i, ok, f.Entry(), pc, f.Name(), file, line)
		}
	}
}