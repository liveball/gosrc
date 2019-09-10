#include "textflag.h"

DATA  foo<>+0x00(SB)/8, $"hello"
DATA  foo<>+0x08(SB)/8, $"go,asm"
DATA  foo<>+0x16(SB)/1, $0x0a
GLOBL foo<>(SB), RODATA, $24

// Syscalls on MacOS are called adding the syscall number to 0x2000000, for example,
// the exit syscall would be 0x2000001

//HelloWorld()
TEXT ·HelloWorld(SB), NOSPLIT, $0
	MOVL 	$(0x2000000+4), AX 	// syscall write
	MOVQ 	$1, DI 			// arg 1 fd 标准输出
	LEAQ 	foo<>(SB), SI 	// arg 2 buf
	MOVL 	$24, DX 		// arg 3 count
	SYSCALL
	RET


//func Neg(x uint64) int64
TEXT ·Neg(SB), NOSPLIT, $0
      MOVQ x+0(FP),AX
      NEGQ AX
      MOVQ AX,ret+8(FP)
      RET
