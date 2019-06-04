package main

import (
	"sync/atomic"
	"unsafe"
)

type g struct {
	param        unsafe.Pointer // passed parameter on wakeup
	atomicstatus uint32
	stackLock    uint32 // sigprof/scang lock; TODO: fold in to atomicstatus
	goid         int64
	schedlink    guintptr
}

type pollDesc struct {
	link *pollDesc
	seq  uintptr // protects from stale timers and ready notifications
	rg   uintptr // pdReady, pdWait, G waiting for read or nil
	wg   uintptr // pdReady, pdWait, G waiting for write or nil
}

type guintptr uintptr

//go:nosplit
func (gp guintptr) ptr() *g { return (*g)(unsafe.Pointer(gp)) }

//go:nosplit
func (gp *guintptr) set(g *g) { *gp = guintptr(unsafe.Pointer(g)) }

//go:nosplit
func (gp *guintptr) cas(old, new guintptr) bool {
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(*gp)), unsafe.Pointer(old), unsafe.Pointer(new))
}

func main() {
	// var i int
	n := int32(100)

	var data [8]byte // to match amd64
	data = [8]byte{'1', '2', '3', '4', '5', '6', '7', '8'}

	var gp guintptr
	for i := int32(0); i < n; i++ {
		var mode int32
		if i%5 == 0 {
			mode += 'r'
		}
		if i%7 == 0 {
			mode += 'w'
		}
		if mode != 0 {
			pd := *(**pollDesc)(unsafe.Pointer(&data))

			netpollready(&gp, pd, mode)
		}
	}
}

func netpollready(gpp *guintptr, pd *pollDesc, mode int32) {
	var rg, wg guintptr
	if mode == 'r' || mode == 'r'+'w' {
		rg.set(netpollunblock(pd, 'r', true))
	}
	if mode == 'w' || mode == 'r'+'w' {
		wg.set(netpollunblock(pd, 'w', true))
	}
	if rg != 0 {
		rg.ptr().schedlink = *gpp
		*gpp = rg
	}
	if wg != 0 {
		wg.ptr().schedlink = *gpp
		*gpp = wg
	}
}

const (
	pdReady uintptr = 1
	pdWait  uintptr = 2
)

func netpollunblock(pd *pollDesc, mode int32, ioready bool) *g {
	gpp := &pd.rg
	if mode == 'w' {
		gpp = &pd.wg
	}

	for {
		old := *gpp
		if old == pdReady {
			return nil
		}
		if old == 0 && !ioready {
			// Only set READY for ioready. runtime_pollWait
			// will check for timeout/cancel before waiting.
			return nil
		}
		var new uintptr
		if ioready {
			new = pdReady
		}
		if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(*gpp)), unsafe.Pointer(old), unsafe.Pointer(new)) { //Casuintptr(gpp, old, new)
			if old == pdReady || old == pdWait {
				old = 0
			}
			return (*g)(unsafe.Pointer(old))
		}
	}
}
