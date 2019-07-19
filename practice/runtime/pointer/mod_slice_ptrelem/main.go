package main

import(
	"fmt"
	"reflect"
	"unsafe"
)

type oper struct {
	rank int
}

var (
	opers []*oper
)

func init() {
	opers = make([]*oper, 0, 3)
	for i := 0; i < 3; i++ {
		opers = append(opers, &oper{rank: i})
	}
}

func main() {
	opers2 := make([]*oper, 0, 3)

	//for _, v := range opers {
	//	if v.rank == 1 {
	//		v.rank = 100
	//	}
	//	opers2 = append(opers2, v)
	//}


	for _, v := range opers {
		vv := new(oper)
		*vv = *v
		if v.rank == 1 {
			vv.rank = 100
		}
		opers2 = append(opers2, vv)
	}



	for _, v := range opers {
		fmt.Printf("opers rank(%d), ptr(%p)\n", v.rank, &v)
	}

	println("------------------")
	for _, v := range opers2 {
		fmt.Printf("opers2 rank(%d), ptr(%p)\n", v.rank, &v)
	}

	fmt.Printf("\n opers:%#v, opers2:%#v\n",
		(*reflect.SliceHeader)(unsafe.Pointer(&opers)),
		(*reflect.SliceHeader)(unsafe.Pointer(&opers2)),
	)
}

//go tool compile -N -l -S main.go | grep "main.go:29"
//0x0131 00305 (main.go:29)       TESTB   AL, (CX)
//0x0133 00307 (main.go:29)       PCDATA  $2, $0
//0x0133 00307 (main.go:29)       MOVQ    $100, (CX)
//0x013a 00314 (main.go:29)       JMP     316

//go tool compile -N -l -S main.go | grep "main.go:39"
//0x0166 00358 (main.go:39)       PCDATA  $2, $4
//0x0166 00358 (main.go:39)       MOVQ    "".vv+136(SP), CX
//0x016e 00366 (main.go:39)       TESTB   AL, (CX)
//0x0170 00368 (main.go:39)       PCDATA  $2, $0
//0x0170 00368 (main.go:39)       MOVQ    $100, (CX)
//0x0177 00375 (main.go:39)       JMP     377
