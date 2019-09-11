//+build amd64,!gccgo
#include "go_asm.h"
#include "textflag.h"

#define INVOKE_SYSCALL	INT	$0x80

#define SYS_sched_getaffinity	242
// func Read(fd int32, p unsafe.Pointer, n int32) int32
TEXT ·Read(SB),NOSPLIT,$0-28
	MOVL	fd+0(FP), DI
	MOVQ	p+8(FP), SI
	MOVL	n+16(FP), DX
	MOVL	$SYS_read, AX
	SYSCALL
	CMPQ	AX, $0xfffffffffffff001
	JLS	2(PC)
	MOVL	$-1, AX
	MOVL	AX, ret+24(FP)
	RET