package main

type oper struct {
	rank int
}

func main() {
	opers := make([]*oper, 0, 3)
	for i := 0; i < 3; i++ {
		opers = append(opers, &oper{rank: i})
	}

	opers2 := make([]*oper, 0, 3)
	for _, v := range opers {
		vv := new(oper)
		*vv = *v
		if v.rank == 1 {
			vv.rank = 100
		}
		opers2 = append(opers2, vv)
	}
}

// go tool compile -N -l -S  demo/type/pointer/slice_ptr/main.go | grep 'demo/type/pointer/slice_ptr/main.go:15'
// 	0x023e 00574 (demo/type/pointer/slice_ptr/main.go:15)	PCDATA	$2, $1
// 	0x023e 00574 (demo/type/pointer/slice_ptr/main.go:15)	LEAQ	type."".oper(SB), AX
// 	0x0245 00581 (demo/type/pointer/slice_ptr/main.go:15)	PCDATA	$2, $0
// 	0x0245 00581 (demo/type/pointer/slice_ptr/main.go:15)	MOVQ	AX, (SP)
// 	0x0249 00585 (demo/type/pointer/slice_ptr/main.go:15)	CALL	runtime.newobject(SB)
// 	0x024e 00590 (demo/type/pointer/slice_ptr/main.go:15)	PCDATA	$2, $1
// 	0x024e 00590 (demo/type/pointer/slice_ptr/main.go:15)	MOVQ	8(SP), AX
// 	0x0253 00595 (demo/type/pointer/slice_ptr/main.go:15)	PCDATA	$0, $11
// 	0x0253 00595 (demo/type/pointer/slice_ptr/main.go:15)	MOVQ	AX, "".vv+96(SP)

// go tool compile -N -l -S  demo/type/pointer/slice_ptr/main.go | grep 'demo/type/pointer/slice_ptr/main.go:16'
// 	0x0258 00600 (demo/type/pointer/slice_ptr/main.go:16)	PCDATA	$2, $6
// 	0x0258 00600 (demo/type/pointer/slice_ptr/main.go:16)	MOVQ	"".v+104(SP), CX
// 	0x025d 00605 (demo/type/pointer/slice_ptr/main.go:16)	TESTB	AL, (CX)
// 	0x025f 00607 (demo/type/pointer/slice_ptr/main.go:16)	PCDATA	$2, $1
// 	0x025f 00607 (demo/type/pointer/slice_ptr/main.go:16)	MOVQ	(CX), CX
// 	0x0262 00610 (demo/type/pointer/slice_ptr/main.go:16)	PCDATA	$2, $0
// 	0x0262 00610 (demo/type/pointer/slice_ptr/main.go:16)	MOVQ	CX, (AX)
