#include "textflag.h"
// 参数大小 = 8 + 4 + 4 , + 4 (默认的 ret符号?)
TEXT ·SwapInt32(SB),NOSPLIT,$0-20
    JMP    ·SwapUint32(SB)
TEXT ·SwapUint32(SB),NOSPLIT,$0-20
    // 第一个参数 移动 8 byte 到 BP
    MOVQ    addr+0(FP), BP
    // 第二个参数 移动 4 byte 到 AX
    MOVL    new+8(FP), AX
    // 原子操作, write-after-read, 把 (AX, offset=0) 与 (BP, offset=0) 交换 4 byte 数据
    XCHGL    AX, 0(BP)
    // 移动 AX 到 old 符号
    MOVL    AX, old+16(FP)
    RET