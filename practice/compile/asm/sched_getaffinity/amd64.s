//+build amd64,!gccgo
#include "go_asm.h"
#include "textflag.h"

#define INVOKE_SYSCALL	INT	$0x80
#define SYS_sched_getaffinity	242
//func sched_getaffinity(pid, len uintptr, buf *byte) int32
TEXT ·sched_getaffinity(SB),NOSPLIT,$0
	MOVL	$SYS_sched_getaffinity, AX
	MOVL	pid+0(FP), BX
	MOVL	len+4(FP), CX
	MOVL	buf+8(FP), DX
	INVOKE_SYSCALL
	INT	$0x80
	MOVL	AX, ret+12(FP)
	RET
