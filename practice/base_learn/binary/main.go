package main

import "fmt"

func main() {
	var d uint8 = 2
	fmt.Printf("%08b [D]\n", d)      // 00000010
	fmt.Printf("%08b (NOT D)\n", ^d) // 11111101

	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)

	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)

	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))

	//setBit()
	//getBit()

	bitOffset := uint8(2)
	on := uint8(1)
	attr := uint32(11)
	fmt.Printf("attr: %08b\n", attr)

	attr2 := AttrSet(attr, bitOffset, on)
	fmt.Printf("attr2: %08b\n", attr2)

	fmt.Println(AttrVal(attr2, bitOffset))
}

func AttrSet(dest uint32, bitOffset, on uint8) (res uint32) {
	fmt.Printf("1 << bitOffset: %08b\n", 1 << bitOffset)
	fmt.Printf("^(1 << bitOffset): %08b\n", ^(1 << bitOffset))
	fmt.Printf("dest: %08b\n", dest)
	fmt.Printf("dest&(^(1 << bitOffset)): %08b\n", dest&(^(1 << bitOffset)))
	fmt.Printf("uint32((on&0x1)<<bitOffset): %08b\n", uint32((on&0x1)<<bitOffset))

	res = dest&(^(1 << bitOffset)) | uint32((on&0x1)<<bitOffset)

	fmt.Printf("res: %08b\n", res)
	return
}

func AttrVal(attr uint32, bitOffset uint8) (v uint8) {
	fmt.Printf("attr: %08b bit:%d\n", attr, bitOffset)
	fmt.Printf("attr>>: %08b bit:%d\n", attr>>bitOffset, (attr>>bitOffset)&1)
	return uint8((attr >> bitOffset) & 1)
}

func setBit() {
	var (
		byte uint

		bitoffset       uint = 0
		bitval, byteval uint
		on              uint = 1
	)

	byte = bitoffset >> 3
	fmt.Printf("byte:%v\n", byte)

	bit := 7 - (bitoffset & 0x7)

	fmt.Printf("bit:%v\n", bit)

	bitval = byteval & (1 << bit)
	fmt.Printf("bitval:%v\n", bitval)

	byteval &= ^(1 << bit)
	byteval |= ((on & 0x1) << bit)

	fmt.Printf("byteval:%v\n", byteval)

}

func getBit() {
	var (
		bitoffset uint = 1
		//bitval    uint
	)

	byte := bitoffset >> 3
	fmt.Printf("byte:%v\n", byte)

	bit := 7 - (bitoffset & 0x7)

	fmt.Printf("bit:%v\n", bit)

	//bitval = ((uint8_t*)o->ptr)[byte] & (1 << bit);
}
