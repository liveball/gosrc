package main

import (
	"fmt"
	"math/rand"
	"unsafe"
)

const (
	wordSize = int(unsafe.Sizeof(uintptr(0)))
	count    = 160 //160
)

var e = make([]byte, count)
var f = make([]byte, count)
var g = make([]byte, count)

func init() {
	for i := 0; i < count; i++ {
		rand.Seed(22)
		e[i] = byte(rand.Intn(256))
		f[i] = byte(rand.Intn(256))
	}
}

func main() {
	fmt.Println(g, e, f)
	safeXORBytes(g, e, f)
	// fastXORWords(g, e, f)
}

func safeXORBytes(dst, a, b []byte) {
	n := len(a)
	ex := n % 8
	// println(ex)//5
	for i := 0; i < ex; i++ {
		dst[i] = a[i] ^ b[i]
		// println(dst[i], a[i], b[i])
	}

	for i := ex; i < n; i += 8 {
		_dst := dst[i : i+8]
		_a := a[i : i+8]
		_b := b[i : i+8]
		_dst[0] = _a[0] ^ _b[0]
		_dst[1] = _a[1] ^ _b[1]
		_dst[2] = _a[2] ^ _b[2]
		_dst[3] = _a[3] ^ _b[3]

		_dst[4] = _a[4] ^ _b[4]
		_dst[5] = _a[5] ^ _b[5]
		_dst[6] = _a[6] ^ _b[6]
		_dst[7] = _a[7] ^ _b[7]
	}
}

func fastXORWords(dst, a, b []byte) {
	dw := *(*[]uintptr)(unsafe.Pointer(&dst))
	aw := *(*[]uintptr)(unsafe.Pointer(&a))
	bw := *(*[]uintptr)(unsafe.Pointer(&b))
	n := len(b) / wordSize
	ex := n % 8
	for i := 0; i < ex; i++ {
		dw[i] = aw[i] ^ bw[i]
	}

	for i := ex; i < n; i += 8 {
		_dw := dw[i : i+8]
		_aw := aw[i : i+8]
		_bw := bw[i : i+8]
		_dw[0] = _aw[0] ^ _bw[0]
		_dw[1] = _aw[1] ^ _bw[1]
		_dw[2] = _aw[2] ^ _bw[2]
		_dw[3] = _aw[3] ^ _bw[3]
		_dw[4] = _aw[4] ^ _bw[4]
		_dw[5] = _aw[5] ^ _bw[5]
		_dw[6] = _aw[6] ^ _bw[6]
		_dw[7] = _aw[7] ^ _bw[7]
	}
}
