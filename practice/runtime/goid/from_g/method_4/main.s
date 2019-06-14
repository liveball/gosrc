#include "textflag.h"

// func getg() interface{}
TEXT ·getg(SB), NOSPLIT, $32-16
    NO_LOCAL_POINTERS

    MOVQ $0, ret_type+0(FP)
    MOVQ $0, ret_data+8(FP)
    GO_RESULTS_INITIALIZED

    // get runtime.g
    MOVQ (TLS), AX

    // get runtime.g type
    MOVQ $type·runtime·g(SB), BX

    // convert (*g) to interface{}
    MOVQ AX, 8(SP)
    MOVQ BX, 0(SP)
    CALL runtime·convT2E(SB)
    MOVQ 16(SP), AX
    MOVQ 24(SP), BX

    // return interface{}
    MOVQ AX, ret_type+0(FP)
    MOVQ BX, ret_data+8(FP)
    RET

    //其中NO_LOCAL_POINTERS表示函数没有局部指针变量。
    //同时对返回的接口进行零值初始化，初始化完成后通过GO_RESULTS_INITIALIZED告知GC。
    //这样可以在保证栈分裂时，GC能够正确处理返回值和局部变量中的指针。

    