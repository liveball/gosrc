#include "textflag.h"

TEXT ·StorepNoWB(SB), NOSPLIT, $0-16
	MOVQ	ptr+0(FP), BX
	MOVQ	val+8(FP), AX
	XCHGQ	AX, 0(BX)
	RET

// bool	·Cas64(uint64 *val, uint64 old, uint64 new)
// Atomically:
//	if(*val == *old){
//		*val = new;
//		return 1;
//	} else {
//		return 0;
//	}
TEXT ·Cas64(SB), NOSPLIT, $0-25
	MOVQ	ptr+0(FP), BX  // BX=val的地址
	MOVQ	old+8(FP), AX  // AX=old
	MOVQ	new+16(FP), CX // CX=new

	//Causes the processor’s LOCK# signal to be asserted during execution of the accompanying instruction (turns the instruction into an atomic instruction).
	//In a multiprocessor environment, the LOCK# signal ensures that the processor has exclusive use of any shared memory while the signal is asserted.
	LOCK

	//Compares the value in the AL, AX, EAX, or RAX register with the first operand (destination operand).
	//If the two values are equal, the second operand (source operand) is loaded into the destination operand.
	//Otherwise, the destination operand is loaded into the AL, AX, EAX or RAX register. RAX register is available only in 64-bit mode.

	//This instruction can be used with a LOCK prefix to allow the instruction to be executed atomically.
	//To simplify the interface to the processor’s bus, the destination operand receives a write cycle without regard to the result of the comparison.
	//The destination operand is written back if the comparison fails; otherwise, the source operand is written into the destination.
	//(The processor never produces a locked read without also producing a locked write.)
	CMPXCHGQ	CX, 0(BX) // 如果AX与BX解引用的值相等，则CX送入BX地址中且ZF置1；否则BX送AX，且ZF清0
	SETEQ	ret+24(FP)
	RET

//func Xadd64(ptr *uint64, delta int64) uint64
TEXT ·Xadd64(SB), NOSPLIT, $0-24
	MOVQ	ptr+0(FP), BX
	MOVQ	delta+8(FP), AX
	MOVQ	AX, CX

	LOCK

    //Exchanges the first operand (destination operand) with the second operand (source operand),
    //then loads the sum of the two values into the destination operand.
    //The destination operand can be a register or a memory location; the source operand is a register.
	XADDQ	AX, 0(BX)
	ADDQ	CX, AX

	MOVQ	AX, ret+16(FP)
	RET
