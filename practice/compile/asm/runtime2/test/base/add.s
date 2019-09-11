#include "textflag.h"

// Add returns the sum of a and b.
//func Add(a int64, b int64) int64
TEXT ·Add(SB), $0-24
    MOVQ a+0(FP), AX
    ADDQ b+8(FP), AX
    MOVQ AX, ret+16(FP)
    RET

