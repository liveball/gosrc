// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Note: some of these functions are semantically inlined
// by the compiler (in src/cmd/compile/internal/gc/ssa.go).

#include "textflag.h"


TEXT runtime∕internal∕atomicx·Cas(SB),NOSPLIT,$0-17
	MOVQ	ptr+0(FP), BX
	MOVL	old+8(FP), AX
	MOVL	new+12(FP), CX
	LOCK
	CMPXCHGL	CX, 0(BX)
	SETEQ	ret+16(FP)
	RET

TEXT runtime∕internal∕atomicx·Cas64(SB), NOSPLIT, $0-25
	MOVQ	ptr+0(FP), BX
	MOVQ	old+8(FP), AX
	MOVQ	new+16(FP), CX
	LOCK
	CMPXCHGQ	CX, 0(BX)
	SETEQ	ret+24(FP)
	RET

TEXT runtime∕internal∕atomicx·Casuintptr(SB), NOSPLIT, $0-25
	JMP	runtime∕internal∕atomicx·Cas64(SB)

// TEXT runtime∕internal∕atomicx·Loaduintptr2(SB), NOSPLIT, $0-16
// 	JMP	runtime∕internal∕atomicx·Load64(SB)
