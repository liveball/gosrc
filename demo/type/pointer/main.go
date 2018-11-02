package main

import (
	"fmt"
	"unsafe"
)

type arc struct {
	aid        uintptr
	attr       *byte
	align      uint8 // alignment of variable with this type
	fieldAlign uint8 // alignment of struct field with this type
}

// ptrType represents a pointer type.
type ptrArc struct {
	arc
	elem *arc
}

func main() {
	var i interface{} = (*unsafe.Pointer)(nil)
	j := *(**ptrArc)(unsafe.Pointer(&i))
	k := *j

	fmt.Printf("i (%#v)\nj(%#v)\nk(%#v)\n\n",
		i, j.arc, k.elem,
	)

	var iptr interface{} = (*unsafe.Pointer)(nil)
	prototype := *(**ptrType)(unsafe.Pointer(&iptr))
	pp := *prototype

	fmt.Printf("pp(%+v)\npp.elem(%+v)\n", pp, pp.elem)
}

type tflag uint8

// a copy of runtime.typeAlg
type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype

type rtype struct {
	size       uintptr
	ptrdata    uintptr  // number of bytes in the type that can contain pointers
	hash       uint32   // hash of type; avoids computation in hash tables
	tflag      tflag    // extra type information flags
	align      uint8    // alignment of variable with this type
	fieldAlign uint8    // alignment of struct field with this type
	kind       uint8    // enumeration for C
	alg        *typeAlg // algorithm table
	gcdata     *byte    // garbage collection data
	str        nameOff  // string form
	ptrToThis  typeOff  // type for pointer to this type, may be zero
}

// ptrType represents a pointer type.
type ptrType struct {
	rtype `reflect:"ptr"`
	elem  *rtype // pointer element (pointed at) type
}
