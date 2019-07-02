package main

import (
	"fmt"
	"runtime"
	"syscall"
	"time"
)

var pid, tid int

// sysNanosleep is defined by OS-specific files (such as runtime_linux_test.go)
// to sleep for the given duration. If nil, dependent tests are skipped.
// The implementation should invoke a blocking system call and not
// call time.Sleep, which would deschedule the goroutine.
var sysNanosleep func(d time.Duration)

func init() {
	// Record pid and tid of init thread for use during test.
	// The call to LockOSThread is just to exercise it;
	// we can't test that it does anything.
	// Instead we're testing that the conditions are good
	// for how it is used in init (must be on main thread).

	// pid, tid = syscall.Getpid(), syscall.Gettid()//for linux
	pid = syscall.Getpid()

	runtime.LockOSThread()

	sysNanosleep = func(d time.Duration) {
		// Invoke a blocking syscall directly; calling time.Sleep()
		// would deschedule the goroutine instead.
		ts := syscall.NsecToTimespec(d.Nanoseconds())
		fmt.Println(ts)
		// for {
		// 	if err := syscall.Nanosleep(&ts, &ts); err != syscall.EINTR {
		// 		return
		// 	}
		// }
	}
}

func main() {
	println(pid, tid)
}
