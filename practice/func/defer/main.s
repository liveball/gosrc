"".foo STEXT size=137 args=0x8 locals=0x28
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	TEXT	"".foo(SB), $40-8 //40 是栈帧大小（局部变量+可能需要的额外调用函数的参数空间的总大小），8是参数及返回值大小(可选的)
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	MOVQ	(TLS), CX//调度相关指令
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	CMPQ	SP, 16(CX)
	0x000d 00013 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	JLS	127
	0x000f 00015 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	SUBQ	$40, SP  //对sp做减法，分配函数栈帧
	0x0013 00019 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	MOVQ	BP, 32(SP)//移动栈
	0x0018 00024 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	LEAQ	32(SP), BP//取地址
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)//gc 相关
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	PCDATA	$2, $0 //gc 相关
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	PCDATA	$0, $0
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	MOVQ	$0, "".r+48(SP) //r=0
	0x0026 00038 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:10)	MOVQ	$5, "".t+24(SP)//t=0
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:13)	PCDATA	$2, $1
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:13)	LEAQ	"".t+24(SP), AX//取t的地址
	0x0034 00052 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:13)	PCDATA	$2, $0
	0x0034 00052 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:13)	MOVQ	AX, 16(SP)
	0x0039 00057 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	MOVL	$8, (SP)
	0x0040 00064 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	PCDATA	$2, $1
	0x0040 00064 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	LEAQ	"".foo.func1·f(SB), AX//取foo 函数的地址
	0x0047 00071 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	PCDATA	$2, $0
	0x0047 00071 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	MOVQ	AX, 8(SP)
	0x004c 00076 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	CALL	runtime.deferproc(SB)//调用运行时 deferproc
	0x0051 00081 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	TESTL	AX, AX
	0x0053 00083 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	JNE	111
	0x0055 00085 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	MOVQ	"".t+24(SP), AX//取t的值
	0x005a 00090 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	MOVQ	AX, "".r+48(SP)
	0x005f 00095 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	XCHGL	AX, AX
	0x0060 00096 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	CALL	runtime.deferreturn(SB)//调用运行时 deferreturn
	0x0065 00101 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	MOVQ	32(SP), BP
	0x006a 00106 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	ADDQ	$40, SP
	0x006e 00110 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:14)	RET
	0x006f 00111 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	XCHGL	AX, AX
	0x0070 00112 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	CALL	runtime.deferreturn(SB) //调用运行时 deferreturn
	0x0075 00117 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	MOVQ	32(SP), BP
	0x007a 00122 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	ADDQ	$40, SP  //对sp做加法，清除函数栈帧
	0x007e 00126 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	RET  
	0x007f 00127 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	NOP
	0x007f 00127 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	PCDATA	$0, $-1
	0x007f 00127 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	PCDATA	$2, $-1
	0x007f 00127 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	CALL	runtime.morestack_noctxt(SB)
	0x0084 00132 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:9)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 70 48  eH..%....H;a.vpH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 44  ..(H.l$ H.l$ H.D
	0x0020 24 30 00 00 00 00 48 c7 44 24 18 05 00 00 00 48  $0....H.D$.....H
	0x0030 8d 44 24 18 48 89 44 24 10 c7 04 24 08 00 00 00  .D$.H.D$...$....
	0x0040 48 8d 05 00 00 00 00 48 89 44 24 08 e8 00 00 00  H......H.D$.....
	0x0050 00 85 c0 75 1a 48 8b 44 24 18 48 89 44 24 30 90  ...u.H.D$.H.D$0.
	0x0060 e8 00 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 90  .....H.l$ H..(..
	0x0070 e8 00 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 e8  .....H.l$ H..(..
	0x0080 00 00 00 00 e9 77 ff ff ff                       .....w...
	rel 5+4 t=16 TLS+0
	rel 67+4 t=15 "".foo.func1·f+0
	rel 77+4 t=8 runtime.deferproc+0
	rel 97+4 t=8 runtime.deferreturn+0
	rel 113+4 t=8 runtime.deferreturn+0
	rel 128+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=602 args=0x0 locals=0x88
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	TEXT	"".main(SB), $136-0
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	MOVQ	(TLS), CX
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	LEAQ	-8(SP), AX
	0x000e 00014 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	CMPQ	AX, 16(CX)
	0x0012 00018 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	JLS	592
	0x0018 00024 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	SUBQ	$136, SP
	0x001f 00031 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	MOVQ	BP, 128(SP)
	0x0027 00039 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	LEAQ	128(SP), BP
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	FUNCDATA	$0, gclocals·3e27b3aa6b89137cce48b3379a2a6610(SB)
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	FUNCDATA	$1, gclocals·27fbe74bf6cfe22d1d3f532b4aaf1d93(SB)
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	FUNCDATA	$3, gclocals·0fce00e8862d709ea3202592b5e747df(SB)
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	PCDATA	$2, $0
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	PCDATA	$0, $0
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	CALL	"".foo(SB)
	0x0034 00052 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	MOVQ	(SP), AX
	0x0038 00056 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	MOVQ	AX, ""..autotmp_14+64(SP)
	0x003d 00061 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	CALL	runtime.printlock(SB)
	0x0042 00066 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	MOVQ	""..autotmp_14+64(SP), AX
	0x0047 00071 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	MOVQ	AX, (SP)
	0x004b 00075 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	CALL	runtime.printint(SB)
	0x0050 00080 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	CALL	runtime.printnl(SB)
	0x0055 00085 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:18)	CALL	runtime.printunlock(SB)
	0x005a 00090 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	MOVQ	$10, 16(SP)
	0x0063 00099 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	MOVL	$8, (SP)
	0x006a 00106 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	PCDATA	$2, $1
	0x006a 00106 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	LEAQ	"".main.func1·f(SB), AX
	0x0071 00113 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	PCDATA	$2, $0
	0x0071 00113 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	MOVQ	AX, 8(SP)
	0x0076 00118 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	CALL	runtime.deferproc(SB)
	0x007b 00123 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	TESTL	AX, AX
	0x007d 00125 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	JNE	570
	0x0083 00131 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	MOVQ	$10, 16(SP)
	0x008c 00140 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	MOVL	$8, (SP)
	0x0093 00147 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	PCDATA	$2, $1
	0x0093 00147 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	LEAQ	"".main.func2·f(SB), AX
	0x009a 00154 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	PCDATA	$2, $0
	0x009a 00154 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	MOVQ	AX, 8(SP)
	0x009f 00159 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	CALL	runtime.deferproc(SB)
	0x00a4 00164 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	TESTL	AX, AX
	0x00a6 00166 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	JNE	548
	0x00ac 00172 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$2, $1
	0x00ac 00172 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	LEAQ	type."".domain(SB), AX
	0x00b3 00179 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$2, $0
	0x00b3 00179 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	MOVQ	AX, (SP)
	0x00b7 00183 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	CALL	runtime.newobject(SB)
	0x00bc 00188 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$2, $2
	0x00bc 00188 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	MOVQ	8(SP), DI
	0x00c1 00193 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$0, $1
	0x00c1 00193 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	MOVQ	DI, "".&a+72(SP)
	0x00c6 00198 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	MOVQ	$2, 8(DI)
	0x00ce 00206 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$2, $-2
	0x00ce 00206 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$0, $-2
	0x00ce 00206 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	CMPL	runtime.writeBarrier(SB), $0
	0x00d5 00213 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	JNE	531
	0x00db 00219 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	LEAQ	go.string."aa"(SB), AX
	0x00e2 00226 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	MOVQ	AX, (DI)
	0x00e5 00229 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x00e5 00229 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$0, $2
	0x00e5 00229 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	AX, ""..autotmp_8+80(SP)
	0x00ea 00234 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	$2, ""..autotmp_8+88(SP)
	0x00f3 00243 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$0, $3
	0x00f3 00243 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	XORPS	X0, X0
	0x00f6 00246 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVUPS	X0, ""..autotmp_7+96(SP)
	0x00fb 00251 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVUPS	X0, ""..autotmp_7+112(SP)
	0x0100 00256 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $1
	0x0100 00256 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	LEAQ	type."".domain(SB), AX
	0x0107 00263 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x0107 00263 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	AX, (SP)
	0x010b 00267 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $1
	0x010b 00267 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	LEAQ	""..autotmp_8+80(SP), AX
	0x0110 00272 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x0110 00272 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	AX, 8(SP)
	0x0115 00277 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	CALL	runtime.convT2E(SB)
	0x011a 00282 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $1
	0x011a 00282 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	24(SP), AX
	0x011f 00287 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	16(SP), CX
	0x0124 00292 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	CX, ""..autotmp_7+96(SP)
	0x0129 00297 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x0129 00297 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	AX, ""..autotmp_7+104(SP)
	0x012e 00302 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $1
	0x012e 00302 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	LEAQ	type.*"".domain(SB), AX
	0x0135 00309 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x0135 00309 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	AX, ""..autotmp_7+112(SP)
	0x013a 00314 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $1
	0x013a 00314 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	"".&a+72(SP), AX
	0x013f 00319 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x013f 00319 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	AX, ""..autotmp_7+120(SP)
	0x0144 00324 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $3
	0x0144 00324 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	LEAQ	go.string."1 a val(%+v) addr(%p)\n"(SB), CX
	0x014b 00331 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x014b 00331 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	CX, (SP)
	0x014f 00335 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	$22, 8(SP)
	0x0158 00344 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $3
	0x0158 00344 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	LEAQ	""..autotmp_7+96(SP), CX
	0x015d 00349 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	PCDATA	$2, $0
	0x015d 00349 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	CX, 16(SP)
	0x0162 00354 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	$2, 24(SP)
	0x016b 00363 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	MOVQ	$2, 32(SP)
	0x0174 00372 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:32)	CALL	fmt.Printf(SB)
	0x0179 00377 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	PCDATA	$2, $1
	0x0179 00377 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	PCDATA	$0, $1
	0x0179 00377 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	"".&a+72(SP), AX
	0x017e 00382 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	8(AX), CX
	0x0182 00386 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	PCDATA	$2, $4
	0x0182 00386 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	(AX), DX
	0x0185 00389 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	PCDATA	$2, $0
	0x0185 00389 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	DX, 16(SP)
	0x018a 00394 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	CX, 24(SP)
	0x018f 00399 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVL	$16, (SP)
	0x0196 00406 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	PCDATA	$2, $3
	0x0196 00406 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	LEAQ	"".fd·f(SB), CX
	0x019d 00413 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	PCDATA	$2, $0
	0x019d 00413 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	CX, 8(SP)
	0x01a2 00418 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	CALL	runtime.deferproc(SB)
	0x01a7 00423 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	TESTL	AX, AX
	0x01a9 00425 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	JNE	509
	0x01ab 00427 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	PCDATA	$2, $1
	0x01ab 00427 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	PCDATA	$0, $0
	0x01ab 00427 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	MOVQ	"".&a+72(SP), AX
	0x01b0 00432 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	PCDATA	$2, $0
	0x01b0 00432 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	MOVQ	AX, 16(SP)
	0x01b5 00437 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	MOVL	$8, (SP)
	0x01bc 00444 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	PCDATA	$2, $1
	0x01bc 00444 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	LEAQ	"".main.func3·f(SB), AX
	0x01c3 00451 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	PCDATA	$2, $0
	0x01c3 00451 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	MOVQ	AX, 8(SP)
	0x01c8 00456 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	CALL	runtime.deferproc(SB)
	0x01cd 00461 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	TESTL	AX, AX
	0x01cf 00463 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	JNE	487
	0x01d1 00465 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:50)	XCHGL	AX, AX
	0x01d2 00466 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:50)	CALL	runtime.deferreturn(SB)
	0x01d7 00471 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:50)	MOVQ	128(SP), BP
	0x01df 00479 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:50)	ADDQ	$136, SP
	0x01e6 00486 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:50)	RET
	0x01e7 00487 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	XCHGL	AX, AX
	0x01e8 00488 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	CALL	runtime.deferreturn(SB)
	0x01ed 00493 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	MOVQ	128(SP), BP
	0x01f5 00501 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	ADDQ	$136, SP
	0x01fc 00508 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	RET
	0x01fd 00509 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	XCHGL	AX, AX
	0x01fe 00510 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	CALL	runtime.deferreturn(SB)
	0x0203 00515 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	MOVQ	128(SP), BP
	0x020b 00523 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	ADDQ	$136, SP
	0x0212 00530 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:33)	RET
	0x0213 00531 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$2, $-2
	0x0213 00531 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	PCDATA	$0, $-2
	0x0213 00531 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	LEAQ	go.string."aa"(SB), AX
	0x021a 00538 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	CALL	runtime.gcWriteBarrier(SB)
	0x021f 00543 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:29)	JMP	229
	0x0224 00548 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	PCDATA	$2, $0
	0x0224 00548 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	PCDATA	$0, $0
	0x0224 00548 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	XCHGL	AX, AX
	0x0225 00549 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	CALL	runtime.deferreturn(SB)
	0x022a 00554 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	MOVQ	128(SP), BP
	0x0232 00562 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	ADDQ	$136, SP
	0x0239 00569 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	RET
	0x023a 00570 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	XCHGL	AX, AX
	0x023b 00571 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	CALL	runtime.deferreturn(SB)
	0x0240 00576 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	MOVQ	128(SP), BP
	0x0248 00584 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	ADDQ	$136, SP
	0x024f 00591 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	RET
	0x0250 00592 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	NOP
	0x0250 00592 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	PCDATA	$0, $-1
	0x0250 00592 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	PCDATA	$2, $-1
	0x0250 00592 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	CALL	runtime.morestack_noctxt(SB)
	0x0255 00597 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:17)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 f8 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 38 02 00 00 48 81 ec 88 00 00 00 48  A...8...H......H
	0x0020 89 ac 24 80 00 00 00 48 8d ac 24 80 00 00 00 e8  ..$....H..$.....
	0x0030 00 00 00 00 48 8b 04 24 48 89 44 24 40 e8 00 00  ....H..$H.D$@...
	0x0040 00 00 48 8b 44 24 40 48 89 04 24 e8 00 00 00 00  ..H.D$@H..$.....
	0x0050 e8 00 00 00 00 e8 00 00 00 00 48 c7 44 24 10 0a  ..........H.D$..
	0x0060 00 00 00 c7 04 24 08 00 00 00 48 8d 05 00 00 00  .....$....H.....
	0x0070 00 48 89 44 24 08 e8 00 00 00 00 85 c0 0f 85 b7  .H.D$...........
	0x0080 01 00 00 48 c7 44 24 10 0a 00 00 00 c7 04 24 08  ...H.D$.......$.
	0x0090 00 00 00 48 8d 05 00 00 00 00 48 89 44 24 08 e8  ...H......H.D$..
	0x00a0 00 00 00 00 85 c0 0f 85 78 01 00 00 48 8d 05 00  ........x...H...
	0x00b0 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b 7c 24  ...H..$.....H.|$
	0x00c0 08 48 89 7c 24 48 48 c7 47 08 02 00 00 00 83 3d  .H.|$HH.G......=
	0x00d0 00 00 00 00 00 0f 85 38 01 00 00 48 8d 05 00 00  .......8...H....
	0x00e0 00 00 48 89 07 48 89 44 24 50 48 c7 44 24 58 02  ..H..H.D$PH.D$X.
	0x00f0 00 00 00 0f 57 c0 0f 11 44 24 60 0f 11 44 24 70  ....W...D$`..D$p
	0x0100 48 8d 05 00 00 00 00 48 89 04 24 48 8d 44 24 50  H......H..$H.D$P
	0x0110 48 89 44 24 08 e8 00 00 00 00 48 8b 44 24 18 48  H.D$......H.D$.H
	0x0120 8b 4c 24 10 48 89 4c 24 60 48 89 44 24 68 48 8d  .L$.H.L$`H.D$hH.
	0x0130 05 00 00 00 00 48 89 44 24 70 48 8b 44 24 48 48  .....H.D$pH.D$HH
	0x0140 89 44 24 78 48 8d 0d 00 00 00 00 48 89 0c 24 48  .D$xH......H..$H
	0x0150 c7 44 24 08 16 00 00 00 48 8d 4c 24 60 48 89 4c  .D$.....H.L$`H.L
	0x0160 24 10 48 c7 44 24 18 02 00 00 00 48 c7 44 24 20  $.H.D$.....H.D$ 
	0x0170 02 00 00 00 e8 00 00 00 00 48 8b 44 24 48 48 8b  .........H.D$HH.
	0x0180 48 08 48 8b 10 48 89 54 24 10 48 89 4c 24 18 c7  H.H..H.T$.H.L$..
	0x0190 04 24 10 00 00 00 48 8d 0d 00 00 00 00 48 89 4c  .$....H......H.L
	0x01a0 24 08 e8 00 00 00 00 85 c0 75 52 48 8b 44 24 48  $........uRH.D$H
	0x01b0 48 89 44 24 10 c7 04 24 08 00 00 00 48 8d 05 00  H.D$...$....H...
	0x01c0 00 00 00 48 89 44 24 08 e8 00 00 00 00 85 c0 75  ...H.D$........u
	0x01d0 16 90 e8 00 00 00 00 48 8b ac 24 80 00 00 00 48  .......H..$....H
	0x01e0 81 c4 88 00 00 00 c3 90 e8 00 00 00 00 48 8b ac  .............H..
	0x01f0 24 80 00 00 00 48 81 c4 88 00 00 00 c3 90 e8 00  $....H..........
	0x0200 00 00 00 48 8b ac 24 80 00 00 00 48 81 c4 88 00  ...H..$....H....
	0x0210 00 00 c3 48 8d 05 00 00 00 00 e8 00 00 00 00 e9  ...H............
	0x0220 c1 fe ff ff 90 e8 00 00 00 00 48 8b ac 24 80 00  ..........H..$..
	0x0230 00 00 48 81 c4 88 00 00 00 c3 90 e8 00 00 00 00  ..H.............
	0x0240 48 8b ac 24 80 00 00 00 48 81 c4 88 00 00 00 c3  H..$....H.......
	0x0250 e8 00 00 00 00 e9 a6 fd ff ff                    ..........
	rel 5+4 t=16 TLS+0
	rel 48+4 t=8 "".foo+0
	rel 62+4 t=8 runtime.printlock+0
	rel 76+4 t=8 runtime.printint+0
	rel 81+4 t=8 runtime.printnl+0
	rel 86+4 t=8 runtime.printunlock+0
	rel 109+4 t=15 "".main.func1·f+0
	rel 119+4 t=8 runtime.deferproc+0
	rel 150+4 t=15 "".main.func2·f+0
	rel 160+4 t=8 runtime.deferproc+0
	rel 175+4 t=15 type."".domain+0
	rel 184+4 t=8 runtime.newobject+0
	rel 208+4 t=15 runtime.writeBarrier+-1
	rel 222+4 t=15 go.string."aa"+0
	rel 259+4 t=15 type."".domain+0
	rel 278+4 t=8 runtime.convT2E+0
	rel 305+4 t=15 type.*"".domain+0
	rel 327+4 t=15 go.string."1 a val(%+v) addr(%p)\n"+0
	rel 373+4 t=8 fmt.Printf+0
	rel 409+4 t=15 "".fd·f+0
	rel 419+4 t=8 runtime.deferproc+0
	rel 447+4 t=15 "".main.func3·f+0
	rel 457+4 t=8 runtime.deferproc+0
	rel 467+4 t=8 runtime.deferreturn+0
	rel 489+4 t=8 runtime.deferreturn+0
	rel 511+4 t=8 runtime.deferreturn+0
	rel 534+4 t=15 go.string."aa"+0
	rel 539+4 t=8 runtime.gcWriteBarrier+0
	rel 550+4 t=8 runtime.deferreturn+0
	rel 572+4 t=8 runtime.deferreturn+0
	rel 593+4 t=8 runtime.morestack_noctxt+0
"".fd STEXT size=295 args=0x10 locals=0x80
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	TEXT	"".fd(SB), $128-16
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	(TLS), CX
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	CMPQ	SP, 16(CX)
	0x000d 00013 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	JLS	285
	0x0013 00019 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	SUBQ	$128, SP
	0x001a 00026 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	BP, 120(SP)
	0x001f 00031 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	LEAQ	120(SP), BP
	0x0024 00036 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	FUNCDATA	$0, gclocals·c85bab9b9628e6cc315c6e757cbf5e5b(SB)
	0x0024 00036 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	FUNCDATA	$1, gclocals·a64e81ec19bbd47020ae781821816ea2(SB)
	0x0024 00036 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	FUNCDATA	$3, gclocals·9fd5fe9155af7b0edeb5d9d14b49308d(SB)
	0x0024 00036 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$2, $1
	0x0024 00036 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$0, $0
	0x0024 00036 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	LEAQ	type."".domain(SB), AX
	0x002b 00043 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$2, $0
	0x002b 00043 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	AX, (SP)
	0x002f 00047 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	CALL	runtime.newobject(SB)
	0x0034 00052 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$2, $2
	0x0034 00052 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	8(SP), DI
	0x0039 00057 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$0, $1
	0x0039 00057 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	DI, "".&a+64(SP)
	0x003e 00062 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	"".a+144(SP), AX
	0x0046 00070 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	AX, 8(DI)
	0x004a 00074 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$2, $-2
	0x004a 00074 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$0, $-2
	0x004a 00074 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	CMPL	runtime.writeBarrier(SB), $0
	0x0051 00081 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	JNE	255
	0x0057 00087 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	"".a+136(SP), CX
	0x005f 00095 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	CX, (DI)
	0x0062 00098 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x0062 00098 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$0, $2
	0x0062 00098 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	CX, ""..autotmp_4+72(SP)
	0x0067 00103 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, ""..autotmp_4+80(SP)
	0x006c 00108 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$0, $3
	0x006c 00108 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	XORPS	X0, X0
	0x006f 00111 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVUPS	X0, ""..autotmp_3+88(SP)
	0x0074 00116 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVUPS	X0, ""..autotmp_3+104(SP)
	0x0079 00121 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x0079 00121 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	LEAQ	type."".domain(SB), AX
	0x0080 00128 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x0080 00128 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, (SP)
	0x0084 00132 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x0084 00132 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	LEAQ	""..autotmp_4+72(SP), AX
	0x0089 00137 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x0089 00137 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, 8(SP)
	0x008e 00142 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	CALL	runtime.convT2E(SB)
	0x0093 00147 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x0093 00147 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	24(SP), AX
	0x0098 00152 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	16(SP), CX
	0x009d 00157 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	CX, ""..autotmp_3+88(SP)
	0x00a2 00162 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x00a2 00162 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, ""..autotmp_3+96(SP)
	0x00a7 00167 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x00a7 00167 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	LEAQ	type.*"".domain(SB), AX
	0x00ae 00174 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x00ae 00174 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, ""..autotmp_3+104(SP)
	0x00b3 00179 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x00b3 00179 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$0, $4
	0x00b3 00179 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	"".&a+64(SP), AX
	0x00b8 00184 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x00b8 00184 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, ""..autotmp_3+112(SP)
	0x00bd 00189 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x00bd 00189 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	LEAQ	go.string."fd a val(%+v) addr(%p)\n"(SB), AX
	0x00c4 00196 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x00c4 00196 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, (SP)
	0x00c8 00200 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	$23, 8(SP)
	0x00d1 00209 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $1
	0x00d1 00209 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	LEAQ	""..autotmp_3+88(SP), AX
	0x00d6 00214 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	PCDATA	$2, $0
	0x00d6 00214 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, 16(SP)
	0x00db 00219 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	$2, 24(SP)
	0x00e4 00228 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	$2, 32(SP)
	0x00ed 00237 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	CALL	fmt.Printf(SB)
	0x00f2 00242 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:54)	PCDATA	$0, $6
	0x00f2 00242 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:54)	MOVQ	120(SP), BP
	0x00f7 00247 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:54)	ADDQ	$128, SP
	0x00fe 00254 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:54)	RET
	0x00ff 00255 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$2, $-2
	0x00ff 00255 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$0, $-2
	0x00ff 00255 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	AX, CX
	0x0102 00258 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	MOVQ	"".a+136(SP), AX
	0x010a 00266 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	CALL	runtime.gcWriteBarrier(SB)
	0x010f 00271 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	AX, DX
	0x0112 00274 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	CX, AX
	0x0115 00277 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:53)	MOVQ	DX, CX
	0x0118 00280 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	JMP	98
	0x011d 00285 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	NOP
	0x011d 00285 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$0, $-1
	0x011d 00285 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	PCDATA	$2, $-1
	0x011d 00285 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	CALL	runtime.morestack_noctxt(SB)
	0x0122 00290 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:52)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 0a  eH..%....H;a....
	0x0010 01 00 00 48 81 ec 80 00 00 00 48 89 6c 24 78 48  ...H......H.l$xH
	0x0020 8d 6c 24 78 48 8d 05 00 00 00 00 48 89 04 24 e8  .l$xH......H..$.
	0x0030 00 00 00 00 48 8b 7c 24 08 48 89 7c 24 40 48 8b  ....H.|$.H.|$@H.
	0x0040 84 24 90 00 00 00 48 89 47 08 83 3d 00 00 00 00  .$....H.G..=....
	0x0050 00 0f 85 a8 00 00 00 48 8b 8c 24 88 00 00 00 48  .......H..$....H
	0x0060 89 0f 48 89 4c 24 48 48 89 44 24 50 0f 57 c0 0f  ..H.L$HH.D$P.W..
	0x0070 11 44 24 58 0f 11 44 24 68 48 8d 05 00 00 00 00  .D$X..D$hH......
	0x0080 48 89 04 24 48 8d 44 24 48 48 89 44 24 08 e8 00  H..$H.D$HH.D$...
	0x0090 00 00 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c  ...H.D$.H.L$.H.L
	0x00a0 24 58 48 89 44 24 60 48 8d 05 00 00 00 00 48 89  $XH.D$`H......H.
	0x00b0 44 24 68 48 8b 44 24 40 48 89 44 24 70 48 8d 05  D$hH.D$@H.D$pH..
	0x00c0 00 00 00 00 48 89 04 24 48 c7 44 24 08 17 00 00  ....H..$H.D$....
	0x00d0 00 48 8d 44 24 58 48 89 44 24 10 48 c7 44 24 18  .H.D$XH.D$.H.D$.
	0x00e0 02 00 00 00 48 c7 44 24 20 02 00 00 00 e8 00 00  ....H.D$ .......
	0x00f0 00 00 48 8b 6c 24 78 48 81 c4 80 00 00 00 c3 48  ..H.l$xH.......H
	0x0100 89 c1 48 8b 84 24 88 00 00 00 e8 00 00 00 00 48  ..H..$.........H
	0x0110 89 c2 48 89 c8 48 89 d1 e9 45 ff ff ff e8 00 00  ..H..H...E......
	0x0120 00 00 e9 d9 fe ff ff                             .......
	rel 5+4 t=16 TLS+0
	rel 39+4 t=15 type."".domain+0
	rel 48+4 t=8 runtime.newobject+0
	rel 76+4 t=15 runtime.writeBarrier+-1
	rel 124+4 t=15 type."".domain+0
	rel 143+4 t=8 runtime.convT2E+0
	rel 170+4 t=15 type.*"".domain+0
	rel 192+4 t=15 go.string."fd a val(%+v) addr(%p)\n"+0
	rel 238+4 t=8 fmt.Printf+0
	rel 267+4 t=8 runtime.gcWriteBarrier+0
	rel 286+4 t=8 runtime.morestack_noctxt+0
"".foo.func1 STEXT nosplit size=10 args=0x8 locals=0x0
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	TEXT	"".foo.func1(SB), NOSPLIT, $0-8
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:11)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:12)	PCDATA	$2, $1
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:12)	PCDATA	$0, $1
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:12)	MOVQ	"".&t+8(SP), AX
	0x0005 00005 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:12)	PCDATA	$2, $0
	0x0005 00005 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:12)	ADDQ	$5, (AX)
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:13)	RET
	0x0000 48 8b 44 24 08 48 83 00 05 c3                    H.D$.H....
"".main.func1 STEXT size=75 args=0x8 locals=0x10
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	TEXT	"".main.func1(SB), $16-8
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	MOVQ	(TLS), CX
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	CMPQ	SP, 16(CX)
	0x000d 00013 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	JLS	68
	0x000f 00015 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	SUBQ	$16, SP
	0x0013 00019 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	MOVQ	BP, 8(SP)
	0x0018 00024 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	LEAQ	8(SP), BP
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	PCDATA	$2, $0
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	PCDATA	$0, $0
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	CALL	runtime.printlock(SB)
	0x0022 00034 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	MOVQ	"".i+24(SP), AX
	0x0027 00039 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	MOVQ	AX, (SP)
	0x002b 00043 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	CALL	runtime.printint(SB)
	0x0030 00048 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	CALL	runtime.printnl(SB)
	0x0035 00053 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:22)	CALL	runtime.printunlock(SB)
	0x003a 00058 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	MOVQ	8(SP), BP
	0x003f 00063 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	ADDQ	$16, SP
	0x0043 00067 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	RET
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:23)	NOP
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	PCDATA	$0, $-1
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	PCDATA	$2, $-1
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	CALL	runtime.morestack_noctxt(SB)
	0x0049 00073 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:21)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 35 48  eH..%....H;a.v5H
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 e8 00 00  ...H.l$.H.l$....
	0x0020 00 00 48 8b 44 24 18 48 89 04 24 e8 00 00 00 00  ..H.D$.H..$.....
	0x0030 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 08 48  ..........H.l$.H
	0x0040 83 c4 10 c3 e8 00 00 00 00 eb b5                 ...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=8 runtime.printlock+0
	rel 44+4 t=8 runtime.printint+0
	rel 49+4 t=8 runtime.printnl+0
	rel 54+4 t=8 runtime.printunlock+0
	rel 69+4 t=8 runtime.morestack_noctxt+0
"".main.func2 STEXT size=75 args=0x8 locals=0x10
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	TEXT	"".main.func2(SB), $16-8
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	MOVQ	(TLS), CX
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	CMPQ	SP, 16(CX)
	0x000d 00013 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	JLS	68
	0x000f 00015 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	SUBQ	$16, SP
	0x0013 00019 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	MOVQ	BP, 8(SP)
	0x0018 00024 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	LEAQ	8(SP), BP
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	PCDATA	$2, $0
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	PCDATA	$0, $0
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	CALL	runtime.printlock(SB)
	0x0022 00034 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	MOVQ	"".a+24(SP), AX
	0x0027 00039 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	MOVQ	AX, (SP)
	0x002b 00043 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	CALL	runtime.printint(SB)
	0x0030 00048 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	CALL	runtime.printnl(SB)
	0x0035 00053 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:26)	CALL	runtime.printunlock(SB)
	0x003a 00058 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	MOVQ	8(SP), BP
	0x003f 00063 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	ADDQ	$16, SP
	0x0043 00067 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	RET
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:27)	NOP
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	PCDATA	$0, $-1
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	PCDATA	$2, $-1
	0x0044 00068 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	CALL	runtime.morestack_noctxt(SB)
	0x0049 00073 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:25)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 35 48  eH..%....H;a.v5H
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 e8 00 00  ...H.l$.H.l$....
	0x0020 00 00 48 8b 44 24 18 48 89 04 24 e8 00 00 00 00  ..H.D$.H..$.....
	0x0030 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 08 48  ..........H.l$.H
	0x0040 83 c4 10 c3 e8 00 00 00 00 eb b5                 ...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=8 runtime.printlock+0
	rel 44+4 t=8 runtime.printint+0
	rel 49+4 t=8 runtime.printnl+0
	rel 54+4 t=8 runtime.printunlock+0
	rel 69+4 t=8 runtime.morestack_noctxt+0
"".main.func3 STEXT size=72 args=0x8 locals=0x18
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	TEXT	"".main.func3(SB), $24-8
	0x0000 00000 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	MOVQ	(TLS), CX
	0x0009 00009 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	CMPQ	SP, 16(CX)
	0x000d 00013 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	JLS	65
	0x000f 00015 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	SUBQ	$24, SP
	0x0013 00019 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	MOVQ	BP, 16(SP)
	0x0018 00024 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	LEAQ	16(SP), BP
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	PCDATA	$2, $1
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	PCDATA	$0, $1
	0x001d 00029 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	MOVQ	"".&a+32(SP), AX
	0x0022 00034 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	MOVQ	8(AX), CX
	0x0026 00038 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	MOVQ	(AX), AX
	0x0029 00041 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	PCDATA	$2, $0
	0x0029 00041 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	MOVQ	AX, (SP)
	0x002d 00045 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	MOVQ	CX, 8(SP)
	0x0032 00050 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:37)	CALL	"".fd(SB)
	0x0037 00055 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	MOVQ	16(SP), BP
	0x003c 00060 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	ADDQ	$24, SP
	0x0040 00064 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	RET
	0x0041 00065 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:38)	NOP
	0x0041 00065 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	PCDATA	$0, $-1
	0x0041 00065 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	PCDATA	$2, $-1
	0x0041 00065 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	CALL	runtime.morestack_noctxt(SB)
	0x0046 00070 (/data/app/go/src/gosrc/demo/base_learn/func/defer/main.go:35)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 32 48  eH..%....H;a.v2H
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 8b 44  ...H.l$.H.l$.H.D
	0x0020 24 20 48 8b 48 08 48 8b 00 48 89 04 24 48 89 4c  $ H.H.H..H..$H.L
	0x0030 24 08 e8 00 00 00 00 48 8b 6c 24 10 48 83 c4 18  $......H.l$.H...
	0x0040 c3 e8 00 00 00 00 eb b8                          ........
	rel 5+4 t=16 TLS+0
	rel 51+4 t=8 "".fd+0
	rel 66+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=92 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	85
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0025 00037 (<autogenerated>:1)	JLS	48
	0x0027 00039 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0027 00039 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0027 00039 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002b 00043 (<autogenerated>:1)	ADDQ	$8, SP
	0x002f 00047 (<autogenerated>:1)	RET
	0x0030 00048 (<autogenerated>:1)	JNE	57
	0x0032 00050 (<autogenerated>:1)	PCDATA	$2, $0
	0x0032 00050 (<autogenerated>:1)	PCDATA	$0, $0
	0x0032 00050 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0037 00055 (<autogenerated>:1)	UNDEF
	0x0039 00057 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0040 00064 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x0045 00069 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0050 00080 (<autogenerated>:1)	ADDQ	$8, SP
	0x0054 00084 (<autogenerated>:1)	RET
	0x0055 00085 (<autogenerated>:1)	NOP
	0x0055 00085 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0055 00085 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005a 00090 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 46 48  eH..%....H;a.vFH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 80 f8 01 76 09 48 8b 2c 24 48 83 c4 08 c3  .....v.H.,$H....
	0x0030 75 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 01  u...............
	0x0040 e8 00 00 00 00 c6 05 00 00 00 00 02 48 8b 2c 24  ............H.,$
	0x0050 48 83 c4 08 c3 e8 00 00 00 00 eb a4              H...........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 51+4 t=8 runtime.throwinit+0
	rel 59+4 t=15 "".initdone·+-1
	rel 65+4 t=8 fmt.init+0
	rel 71+4 t=15 "".initdone·+-1
	rel 86+4 t=8 runtime.morestack_noctxt+0
type..hash.[2]interface {} STEXT dupok size=110 args=0x18 locals=0x28
	0x0000 00000 (<autogenerated>:1)	TEXT	type..hash.[2]interface {}(SB), DUPOK, $40-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	103
	0x000f 00015 (<autogenerated>:1)	SUBQ	$40, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 32(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	32(SP), BP
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (<autogenerated>:1)	FUNCDATA	$3, gclocals·ee104e299ed2e4539b82c61c5a4b843d(SB)
	0x001d 00029 (<autogenerated>:1)	PCDATA	$2, $0
	0x001d 00029 (<autogenerated>:1)	PCDATA	$0, $0
	0x001d 00029 (<autogenerated>:1)	XORL	AX, AX
	0x001f 00031 (<autogenerated>:1)	MOVQ	"".h+56(SP), CX
	0x0024 00036 (<autogenerated>:1)	JMP	82
	0x0026 00038 (<autogenerated>:1)	MOVQ	AX, "".i+24(SP)
	0x002b 00043 (<autogenerated>:1)	SHLQ	$4, AX
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $1
	0x002f 00047 (<autogenerated>:1)	MOVQ	"".p+48(SP), BX
	0x0034 00052 (<autogenerated>:1)	PCDATA	$2, $2
	0x0034 00052 (<autogenerated>:1)	ADDQ	BX, AX
	0x0037 00055 (<autogenerated>:1)	PCDATA	$2, $0
	0x0037 00055 (<autogenerated>:1)	MOVQ	AX, (SP)
	0x003b 00059 (<autogenerated>:1)	MOVQ	CX, 8(SP)
	0x0040 00064 (<autogenerated>:1)	CALL	runtime.nilinterhash(SB)
	0x0045 00069 (<autogenerated>:1)	MOVQ	16(SP), CX
	0x004a 00074 (<autogenerated>:1)	MOVQ	"".i+24(SP), AX
	0x004f 00079 (<autogenerated>:1)	INCQ	AX
	0x0052 00082 (<autogenerated>:1)	CMPQ	AX, $2
	0x0056 00086 (<autogenerated>:1)	JLT	38
	0x0058 00088 (<autogenerated>:1)	PCDATA	$0, $1
	0x0058 00088 (<autogenerated>:1)	MOVQ	CX, "".~r2+64(SP)
	0x005d 00093 (<autogenerated>:1)	MOVQ	32(SP), BP
	0x0062 00098 (<autogenerated>:1)	ADDQ	$40, SP
	0x0066 00102 (<autogenerated>:1)	RET
	0x0067 00103 (<autogenerated>:1)	NOP
	0x0067 00103 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0067 00103 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0067 00103 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x006c 00108 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 58 48  eH..%....H;a.vXH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 31 c0 48  ..(H.l$ H.l$ 1.H
	0x0020 8b 4c 24 38 eb 2c 48 89 44 24 18 48 c1 e0 04 48  .L$8.,H.D$.H...H
	0x0030 8b 5c 24 30 48 01 d8 48 89 04 24 48 89 4c 24 08  .\$0H..H..$H.L$.
	0x0040 e8 00 00 00 00 48 8b 4c 24 10 48 8b 44 24 18 48  .....H.L$.H.D$.H
	0x0050 ff c0 48 83 f8 02 7c ce 48 89 4c 24 40 48 8b 6c  ..H...|.H.L$@H.l
	0x0060 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb 92        $ H..(........
	rel 5+4 t=16 TLS+0
	rel 65+4 t=8 runtime.nilinterhash+0
	rel 104+4 t=8 runtime.morestack_noctxt+0
type..eq.[2]interface {} STEXT dupok size=182 args=0x18 locals=0x30
	0x0000 00000 (<autogenerated>:1)	TEXT	type..eq.[2]interface {}(SB), DUPOK, $48-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	172
	0x0013 00019 (<autogenerated>:1)	SUBQ	$48, SP
	0x0017 00023 (<autogenerated>:1)	MOVQ	BP, 40(SP)
	0x001c 00028 (<autogenerated>:1)	LEAQ	40(SP), BP
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (<autogenerated>:1)	FUNCDATA	$3, gclocals·a1bdf42ea3370bf425f59e11a41daee2(SB)
	0x0021 00033 (<autogenerated>:1)	PCDATA	$2, $1
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	MOVQ	"".p+56(SP), AX
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $2
	0x0026 00038 (<autogenerated>:1)	MOVQ	"".q+64(SP), CX
	0x002b 00043 (<autogenerated>:1)	XORL	DX, DX
	0x002d 00045 (<autogenerated>:1)	JMP	72
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $0
	0x002f 00047 (<autogenerated>:1)	MOVQ	""..autotmp_8+32(SP), BX
	0x0034 00052 (<autogenerated>:1)	LEAQ	1(BX), DX
	0x0038 00056 (<autogenerated>:1)	PCDATA	$2, $3
	0x0038 00056 (<autogenerated>:1)	MOVQ	"".p+56(SP), BX
	0x003d 00061 (<autogenerated>:1)	PCDATA	$2, $4
	0x003d 00061 (<autogenerated>:1)	MOVQ	"".q+64(SP), SI
	0x0042 00066 (<autogenerated>:1)	PCDATA	$2, $5
	0x0042 00066 (<autogenerated>:1)	MOVQ	BX, AX
	0x0045 00069 (<autogenerated>:1)	PCDATA	$2, $2
	0x0045 00069 (<autogenerated>:1)	MOVQ	SI, CX
	0x0048 00072 (<autogenerated>:1)	CMPQ	DX, $2
	0x004c 00076 (<autogenerated>:1)	JGE	157
	0x004e 00078 (<autogenerated>:1)	MOVQ	DX, BX
	0x0051 00081 (<autogenerated>:1)	SHLQ	$4, DX
	0x0055 00085 (<autogenerated>:1)	PCDATA	$2, $6
	0x0055 00085 (<autogenerated>:1)	MOVQ	8(DX)(AX*1), SI
	0x005a 00090 (<autogenerated>:1)	PCDATA	$2, $7
	0x005a 00090 (<autogenerated>:1)	MOVQ	(DX)(AX*1), DI
	0x005e 00094 (<autogenerated>:1)	PCDATA	$2, $8
	0x005e 00094 (<autogenerated>:1)	MOVQ	8(DX)(CX*1), R8
	0x0063 00099 (<autogenerated>:1)	PCDATA	$2, $9
	0x0063 00099 (<autogenerated>:1)	MOVQ	(DX)(CX*1), DX
	0x0067 00103 (<autogenerated>:1)	CMPQ	DI, DX
	0x006a 00106 (<autogenerated>:1)	JNE	142
	0x006c 00108 (<autogenerated>:1)	MOVQ	BX, ""..autotmp_8+32(SP)
	0x0071 00113 (<autogenerated>:1)	MOVQ	DI, (SP)
	0x0075 00117 (<autogenerated>:1)	PCDATA	$2, $10
	0x0075 00117 (<autogenerated>:1)	MOVQ	SI, 8(SP)
	0x007a 00122 (<autogenerated>:1)	PCDATA	$2, $0
	0x007a 00122 (<autogenerated>:1)	MOVQ	R8, 16(SP)
	0x007f 00127 (<autogenerated>:1)	CALL	runtime.efaceeq(SB)
	0x0084 00132 (<autogenerated>:1)	PCDATA	$2, $1
	0x0084 00132 (<autogenerated>:1)	LEAQ	24(SP), AX
	0x0089 00137 (<autogenerated>:1)	PCDATA	$2, $0
	0x0089 00137 (<autogenerated>:1)	CMPB	(AX), $0
	0x008c 00140 (<autogenerated>:1)	JNE	47
	0x008e 00142 (<autogenerated>:1)	PCDATA	$0, $1
	0x008e 00142 (<autogenerated>:1)	MOVB	$0, "".~r2+72(SP)
	0x0093 00147 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x0098 00152 (<autogenerated>:1)	ADDQ	$48, SP
	0x009c 00156 (<autogenerated>:1)	RET
	0x009d 00157 (<autogenerated>:1)	MOVB	$1, "".~r2+72(SP)
	0x00a2 00162 (<autogenerated>:1)	MOVQ	40(SP), BP
	0x00a7 00167 (<autogenerated>:1)	ADDQ	$48, SP
	0x00ab 00171 (<autogenerated>:1)	RET
	0x00ac 00172 (<autogenerated>:1)	NOP
	0x00ac 00172 (<autogenerated>:1)	PCDATA	$0, $-1
	0x00ac 00172 (<autogenerated>:1)	PCDATA	$2, $-1
	0x00ac 00172 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x00b1 00177 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 99  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 30 48 89 6c 24 28 48 8d 6c 24  ...H..0H.l$(H.l$
	0x0020 28 48 8b 44 24 38 48 8b 4c 24 40 31 d2 eb 19 48  (H.D$8H.L$@1...H
	0x0030 8b 5c 24 20 48 8d 53 01 48 8b 5c 24 38 48 8b 74  .\$ H.S.H.\$8H.t
	0x0040 24 40 48 89 d8 48 89 f1 48 83 fa 02 7d 4f 48 89  $@H..H..H...}OH.
	0x0050 d3 48 c1 e2 04 48 8b 74 02 08 48 8b 3c 02 4c 8b  .H...H.t..H.<.L.
	0x0060 44 0a 08 48 8b 14 0a 48 39 d7 75 22 48 89 5c 24  D..H...H9.u"H.\$
	0x0070 20 48 89 3c 24 48 89 74 24 08 4c 89 44 24 10 e8   H.<$H.t$.L.D$..
	0x0080 00 00 00 00 48 8d 44 24 18 80 38 00 75 a1 c6 44  ....H.D$..8.u..D
	0x0090 24 48 00 48 8b 6c 24 28 48 83 c4 30 c3 c6 44 24  $H.H.l$(H..0..D$
	0x00a0 48 01 48 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00  H.H.l$(H..0.....
	0x00b0 00 e9 4a ff ff ff                                ..J...
	rel 5+4 t=16 TLS+0
	rel 128+4 t=8 runtime.efaceeq+0
	rel 173+4 t=8 runtime.morestack_noctxt+0
go.loc."".foo SDWARFLOC size=103
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 26 00 00 00 00 00 00 00 89 00 00 00 00 00 00 00  &...............
	0x0020 01 00 9c 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 2f 00 00 00 00 00 00 00 89 00 00 00 00  .../............
	0x0050 00 00 00 02 00 91 68 00 00 00 00 00 00 00 00 00  ......h.........
	0x0060 00 00 00 00 00 00 00                             .......
	rel 8+8 t=1 "".foo+0
	rel 59+8 t=1 "".foo+0
go.info."".foo SDWARFINFO size=57
	0x0000 02 22 22 2e 66 6f 6f 00 00 00 00 00 00 00 00 00  ."".foo.........
	0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f  ................
	0x0020 72 00 01 09 00 00 00 00 00 00 00 00 0a 74 00 0a  r............t..
	0x0030 00 00 00 00 00 00 00 00 00                       .........
	rel 8+8 t=1 "".foo+0
	rel 16+8 t=1 "".foo+137
	rel 26+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 36+4 t=28 go.info.int+0
	rel 40+4 t=28 go.loc."".foo+0
	rel 48+4 t=28 go.info.int+0
	rel 52+4 t=28 go.loc."".foo+51
go.range."".foo SDWARFRANGE size=0
go.isstmt."".foo SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 17 01 05 02 07 01 11 02 02 01 02  ................
	0x0010 02 05 01 1b 02 14 00                             .......
go.string."aa" SRODATA dupok size=2
	0x0000 61 61                                            aa
go.string."1 a val(%+v) addr(%p)\n" SRODATA dupok size=22
	0x0000 31 20 61 20 76 61 6c 28 25 2b 76 29 20 61 64 64  1 a val(%+v) add
	0x0010 72 28 25 70 29 0a                                r(%p).
go.loc."".main SDWARFLOC size=72
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 c1 00 00 00 00 00 00 00 1a 01 00 00 00 00 00 00  ................
	0x0020 01 00 55 1a 01 00 00 00 00 00 00 5a 02 00 00 00  ..U........Z....
	0x0030 00 00 00 03 00 91 b8 7f 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 8+8 t=1 "".main+0
go.info."".main SDWARFINFO size=46
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0a 26 61 00 1d 00 00 00 00 00 00 00 00 00        .&a...........
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+602
	rel 27+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 37+4 t=28 go.info.*"".domain+0
	rel 41+4 t=28 go.loc."".main+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 18 04 17 03 05 01 26 02 10 01 11 02 02 01 06  .......&........
	0x0010 02 10 01 11 02 02 01 06 02 07 01 32 02 05 01 8f  ...........2....
	0x0020 01 02 05 01 29 02 02 01 02 02 05 01 05 02 07 01  ....)...........
	0x0030 11 02 02 01 02 02 01 01 05 02 10 01 06 02 11 01  ................
	0x0040 05 02 10 01 11 02 01 01 05 02 11 01 05 02 1a 00  ................
go.string."fd a val(%+v) addr(%p)\n" SRODATA dupok size=23
	0x0000 66 64 20 61 20 76 61 6c 28 25 2b 76 29 20 61 64  fd a val(%+v) ad
	0x0010 64 72 28 25 70 29 0a                             dr(%p).
go.loc."".fd SDWARFLOC size=72
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 39 00 00 00 00 00 00 00 93 00 00 00 00 00 00 00  9...............
	0x0020 01 00 55 93 00 00 00 00 00 00 00 27 01 00 00 00  ..U........'....
	0x0030 00 00 00 03 00 91 b8 7f 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 8+8 t=1 "".fd+0
go.info."".fd SDWARFINFO size=54
	0x0000 02 22 22 2e 66 64 00 00 00 00 00 00 00 00 00 00  ."".fd..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0a 26  ...............&
	0x0020 61 00 34 00 00 00 00 00 00 00 00 0e 61 00 00 34  a.4.........a..4
	0x0030 00 00 00 00 00 00                                ......
	rel 7+8 t=1 "".fd+0
	rel 15+8 t=1 "".fd+295
	rel 25+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 35+4 t=28 go.info.*"".domain+0
	rel 39+4 t=28 go.loc."".fd+0
	rel 48+4 t=28 go.info."".domain+0
go.range."".fd SDWARFRANGE size=0
go.isstmt."".fd SDWARFMISC size=0
	0x0000 04 13 04 11 03 07 01 37 02 05 01 8b 01 02 0d 01  .......7........
	0x0010 1e 02 0a 00                                      ....
go.loc."".foo.func1 SDWARFLOC size=0
go.info."".foo.func1 SDWARFINFO size=49
	0x0000 02 22 22 2e 66 6f 6f 2e 66 75 6e 63 31 00 00 00  ."".foo.func1...
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 9c  ................
	0x0020 00 00 00 00 01 0e 26 74 00 00 0b 00 00 00 00 00  ......&t........
	0x0030 00                                               .
	rel 14+8 t=1 "".foo.func1+0
	rel 22+8 t=1 "".foo.func1+10
	rel 32+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 43+4 t=28 go.info.*int+0
go.range."".foo.func1 SDWARFRANGE size=0
go.isstmt."".foo.func1 SDWARFMISC size=0
	0x0000 04 05 01 04 02 01 00                             .......
go.loc."".main.func1 SDWARFLOC size=0
go.info."".main.func1 SDWARFINFO size=49
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 31 00 00  ."".main.func1..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0e 69 00 00 16 00 00 00 00 00  .......i........
	0x0030 00                                               .
	rel 15+8 t=1 "".main.func1+0
	rel 23+8 t=1 "".main.func1+75
	rel 33+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 43+4 t=28 go.info.int+0
go.range."".main.func1 SDWARFRANGE size=0
go.isstmt."".main.func1 SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 05 01 18 02 11 00                 ...........
go.loc."".main.func2 SDWARFLOC size=0
go.info."".main.func2 SDWARFINFO size=49
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 32 00 00  ."".main.func2..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0e 61 00 00 19 00 00 00 00 00  .......a........
	0x0030 00                                               .
	rel 15+8 t=1 "".main.func2+0
	rel 23+8 t=1 "".main.func2+75
	rel 33+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 43+4 t=28 go.info.int+0
go.range."".main.func2 SDWARFRANGE size=0
go.isstmt."".main.func2 SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 05 01 18 02 11 00                 ...........
go.loc."".main.func3 SDWARFLOC size=0
go.info."".main.func3 SDWARFINFO size=50
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 33 00 00  ."".main.func3..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0e 26 61 00 00 23 00 00 00 00  .......&a..#....
	0x0030 00 00                                            ..
	rel 15+8 t=1 "".main.func3+0
	rel 23+8 t=1 "".main.func3+72
	rel 33+4 t=29 gofile../data/app/go/src/gosrc/demo/base_learn/func/defer/main.go+0
	rel 44+4 t=28 go.info.*"".domain+0
go.range."".main.func3 SDWARFRANGE size=0
go.isstmt."".main.func3 SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 05 01 15 02 11 00                 ...........
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+92
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 05 02 09 01 07 02 09 01 15  ................
	0x0010 02 07 00                                         ...
"".initdone· SNOPTRBSS size=1
"".fd·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".fd+0
"".foo.func1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".foo.func1+0
"".main.func1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func1+0
"".main.func2·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func2+0
"".main.func3·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func3+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.domain- SRODATA dupok size=15
	0x0000 00 00 0c 2a 6d 61 69 6e 2e 64 6f 6d 61 69 6e     ...*main.domain
type.*"".domain SRODATA size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 24 73 d7 dc 00 08 08 36 00 00 00 00 00 00 00 00  $s.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.domain-+0
	rel 48+8 t=1 type."".domain+0
type..namedata.do- SRODATA dupok size=5
	0x0000 00 00 02 64 6f                                   ...do
type."".domain SRODATA size=120
	0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b eb 60 98 07 08 08 19 00 00 00 00 00 00 00 00  ..`.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+112
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.domain-+0
	rel 44+4 t=5 type.*"".domain+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".domain+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.do-+0
	rel 104+8 t=1 type.string+0
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
go.loc.type..hash.[2]interface {} SDWARFLOC dupok size=173
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 2b 00 00 00 00 00 00 00 56 00 00 00 00 00 00 00  +.......V.......
	0x0020 02 00 91 68 56 00 00 00 00 00 00 00 6e 00 00 00  ...hV.......n...
	0x0030 00 00 00 00 01 00 50 00 00 00 00 00 00 00 00 00  ......P.........
	0x0040 00 00 00 00 00 00 00 ff ff ff ff ff ff ff ff 00  ................
	0x0050 00 00 00 00 00 00 00 1f 00 00 00 00 00 00 00 6e  ...............n
	0x0060 00 00 00 00 00 00 00 01 00 9c 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 ff ff ff ff ff ff  ................
	0x0080 ff ff 00 00 00 00 00 00 00 00 4a 00 00 00 00 00  ..........J.....
	0x0090 00 00 6e 00 00 00 00 00 00 00 01 00 52 00 00 00  ..n.........R...
	0x00a0 00 00 00 00 00 00 00 00 00 00 00 00 00           .............
	rel 8+8 t=1 type..hash.[2]interface {}+0
	rel 79+8 t=1 type..hash.[2]interface {}+0
	rel 130+8 t=1 type..hash.[2]interface {}+0
go.info.type..hash.[2]interface {} SDWARFINFO dupok size=102
	0x0000 02 74 79 70 65 2e 2e 68 61 73 68 2e 5b 32 5d 69  .type..hash.[2]i
	0x0010 6e 74 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00  nterface {}.....
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0030 00 00 01 0a 69 00 01 00 00 00 00 00 00 00 00 0f  ....i...........
	0x0040 70 00 00 01 00 00 00 00 00 00 00 00 0f 68 00 00  p............h..
	0x0050 01 00 00 00 00 00 00 00 00 0e 7e 72 32 00 01 01  ..........~r2...
	0x0060 00 00 00 00 00 00                                ......
	rel 28+8 t=1 type..hash.[2]interface {}+0
	rel 36+8 t=1 type..hash.[2]interface {}+110
	rel 46+4 t=29 gofile..<autogenerated>+0
	rel 55+4 t=28 go.info.int+0
	rel 59+4 t=28 go.loc.type..hash.[2]interface {}+0
	rel 68+4 t=28 go.info.*[2]interface {}+0
	rel 72+4 t=28 go.loc.type..hash.[2]interface {}+71
	rel 81+4 t=28 go.info.uintptr+0
	rel 85+4 t=28 go.loc.type..hash.[2]interface {}+122
	rel 96+4 t=28 go.info.uintptr+0
go.range.type..hash.[2]interface {} SDWARFRANGE dupok size=0
go.isstmt.type..hash.[2]interface {} SDWARFMISC dupok size=0
	0x0000 04 0f 04 0e 03 02 01 33 02 04 01 02 02 05 01 0a  .......3........
	0x0010 02 07 00                                         ...
go.loc.type..eq.[2]interface {} SDWARFLOC dupok size=154
	0x0000 ff ff ff ff ff ff ff ff 00 00 00 00 00 00 00 00  ................
	0x0010 4c 00 00 00 00 00 00 00 55 00 00 00 00 00 00 00  L.......U.......
	0x0020 01 00 51 00 00 00 00 00 00 00 00 00 00 00 00 00  ..Q.............
	0x0030 00 00 00 ff ff ff ff ff ff ff ff 00 00 00 00 00  ................
	0x0040 00 00 00 26 00 00 00 00 00 00 00 b6 00 00 00 00  ...&............
	0x0050 00 00 00 01 00 9c 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 ff ff ff ff ff ff ff ff 00 00  ................
	0x0070 00 00 00 00 00 00 26 00 00 00 00 00 00 00 b6 00  ......&.........
	0x0080 00 00 00 00 00 00 02 00 91 08 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00                    ..........
	rel 8+8 t=1 type..eq.[2]interface {}+0
	rel 59+8 t=1 type..eq.[2]interface {}+0
	rel 110+8 t=1 type..eq.[2]interface {}+0
go.info.type..eq.[2]interface {} SDWARFINFO dupok size=100
	0x0000 02 74 79 70 65 2e 2e 65 71 2e 5b 32 5d 69 6e 74  .type..eq.[2]int
	0x0010 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00 00 00  erface {}.......
	0x0020 00 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00  ................
	0x0030 01 0a 69 00 01 00 00 00 00 00 00 00 00 0f 70 00  ..i...........p.
	0x0040 00 01 00 00 00 00 00 00 00 00 0f 71 00 00 01 00  ...........q....
	0x0050 00 00 00 00 00 00 00 0e 7e 72 32 00 01 01 00 00  ........~r2.....
	0x0060 00 00 00 00                                      ....
	rel 26+8 t=1 type..eq.[2]interface {}+0
	rel 34+8 t=1 type..eq.[2]interface {}+182
	rel 44+4 t=29 gofile..<autogenerated>+0
	rel 53+4 t=28 go.info.int+0
	rel 57+4 t=28 go.loc.type..eq.[2]interface {}+0
	rel 66+4 t=28 go.info.*[2]interface {}+0
	rel 70+4 t=28 go.loc.type..eq.[2]interface {}+51
	rel 79+4 t=28 go.info.*[2]interface {}+0
	rel 83+4 t=28 go.loc.type..eq.[2]interface {}+102
	rel 94+4 t=28 go.info.bool+0
go.range.type..eq.[2]interface {} SDWARFRANGE dupok size=0
go.isstmt.type..eq.[2]interface {} SDWARFMISC dupok size=0
	0x0000 04 13 04 0e 03 05 01 22 02 04 01 42 02 05 01 0a  ......."...B....
	0x0010 02 05 01 0a 02 0a 00                             .......
type..hashfunc.[2]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash.[2]interface {}+0
type..eqfunc.[2]interface {} SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[2]interface {}+0
type..alg.[2]interface {} SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc.[2]interface {}+0
	rel 8+8 t=1 type..eqfunc.[2]interface {}+0
type..namedata.*[2]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
	0x0010 20 7b 7d                                          {}
type.*[2]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 be 73 2d 71 00 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[2]interface {}-+0
	rel 48+8 t=1 type.[2]interface {}+0
runtime.gcbits.0a SRODATA dupok size=1
	0x0000 0a                                               .
type.[2]interface {} SRODATA dupok size=72
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg.[2]interface {}+0
	rel 32+8 t=1 runtime.gcbits.0a+0
	rel 40+4 t=5 type..namedata.*[2]interface {}-+0
	rel 44+4 t=6 type.*[2]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·3e27b3aa6b89137cce48b3379a2a6610 SRODATA dupok size=8
	0x0000 05 00 00 00 00 00 00 00                          ........
gclocals·27fbe74bf6cfe22d1d3f532b4aaf1d93 SRODATA dupok size=13
	0x0000 05 00 00 00 07 00 00 00 00 01 03 53 51           ...........SQ
gclocals·0fce00e8862d709ea3202592b5e747df SRODATA dupok size=13
	0x0000 05 00 00 00 07 00 00 00 00 01 40 02 04           ..........@..
gclocals·c85bab9b9628e6cc315c6e757cbf5e5b SRODATA dupok size=15
	0x0000 07 00 00 00 01 00 00 00 01 01 00 00 00 00 00     ...............
gclocals·a64e81ec19bbd47020ae781821816ea2 SRODATA dupok size=15
	0x0000 07 00 00 00 07 00 00 00 00 01 03 53 52 50 00     ...........SRP.
gclocals·9fd5fe9155af7b0edeb5d9d14b49308d SRODATA dupok size=11
	0x0000 03 00 00 00 07 00 00 00 00 01 40                 ..........@
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·ee104e299ed2e4539b82c61c5a4b843d SRODATA dupok size=11
	0x0000 03 00 00 00 04 00 00 00 00 08 01                 ...........
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
gclocals·a1bdf42ea3370bf425f59e11a41daee2 SRODATA dupok size=19
	0x0000 0b 00 00 00 08 00 00 00 00 01 03 08 28 21 23 22  ............(!#"
	0x0010 a2 a0 80                                         ...
