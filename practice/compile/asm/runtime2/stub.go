package runtime2

import "unsafe"

func getg() unsafe.Pointer

func setg(gg *g)

// StorepNoWB performs *ptr = val atomically and without a write
// barrier.
//
// NO go:noescape annotation; see atomic_pointer.go.
func StorepNoWB(ptr unsafe.Pointer, val unsafe.Pointer)


//go:noescape
func Cas64(ptr *uint64, old, new uint64) bool

//go:noescape
func Xadd64(ptr *uint64, delta int64) uint64