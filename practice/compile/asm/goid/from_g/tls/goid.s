#include "textflag.h"
#include "go_tls.h"

// 返回值 8 bytes, 符号为 getg
//func getg() unsafe.Pointer
TEXT ·getg(SB), NOSPLIT, $0-8
    // get_tls 的宏为： #define    get_tls(r)    MOVQ TLS, r
    // 等价于 MOVQ TLS, CX
    // 从 TLS(Thread Local Storage) 起始移动 8 byte 值 到 CX 寄存器
    get_tls(CX)
    // g的宏为： g(r)    0(r)(TLS*1)
    // 等价于 0(CX)(TLS*1), AX
    // 查到意义为 indexed with offset, 这里 offset=0, 索引是什么意思不清楚
    MOVQ g(CX), AX
    // 从AX起始移动 8 byte 值，到ret符号的位置
    MOVQ AX, ret+0(FP)
    RET
