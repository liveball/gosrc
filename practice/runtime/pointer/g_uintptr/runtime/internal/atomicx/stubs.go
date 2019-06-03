// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !wasm

package atomicx

import "unsafe"

//go:noescape
func Cas(ptr *uint32, old, new uint32) bool

// NO go:noescape annotation; see atomic_pointer.go.
func Casp1(ptr *unsafe.Pointer, old, new unsafe.Pointer) bool

//go:noescape
func Casuintptr(ptr *uintptr, old, new uintptr) bool

//go:noescape
// func Storeuintptr(ptr *uintptr, new uintptr)

//go:noescape
func Loaduintptr2(ptr *uintptr) uintptr
