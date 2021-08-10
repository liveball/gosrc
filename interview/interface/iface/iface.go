package main

import (
	"github.com/davecgh/go-spew/spew"
	"unsafe"
)

type imethod struct {
	name nameOff
	ityp typeOff
}

// name is an encoded type name with optional extra data.
// See reflect/type.go for details.
type name struct {
	bytes *byte
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

type tflag uint8
type nameOff int32
type typeOff int32
type textOff int32

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        nameOff
	ptrToThis  typeOff
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type Person interface {
	Say() string
}

type Man struct {
	Name string
}

func (m *Man) Say() string {
	return "Man"
}

func main() {
	var p Person

	m := &Man{Name: "hhf"}
	p = m
	println(p.Say())


	spew.Dump(p,m)

	spew.Dump("itab:", *(*iface)(unsafe.Pointer(&p)).tab)
	spew.Dump("data:", *(*Man)((*iface)(unsafe.Pointer(&p)).data))
}
