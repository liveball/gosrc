// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Time-related runtime and pieces of package time.

package tls

// Package time knows the layout of this structure.
// If this struct changes, adjust ../time/sleep.go:/runtimeTimer.
// For GOOS=nacl, package syscall knows the layout of this structure.
// If this struct changes, adjust ../syscall/net_nacl.go:/runtimeTimer.
type timer struct {
	tb *timersBucket // the bucket the timer lives in
	i  int           // heap index

	// Timer wakes up at when, and then at when+period, ... (period > 0 only)
	// each time calling f(arg, now) in the timer goroutine, so f must be
	// a well-behaved function and not block.
	when   int64
	period int64
	f      func(interface{}, uintptr)
	arg    interface{}
	seq    uintptr
}

//go:notinheap
type timersBucket struct {
	lock         mutex
	gp           *g
	created      bool
	sleeping     bool
	rescheduling bool
	sleepUntil   int64
	waitnote     note
	t            []*timer
}
