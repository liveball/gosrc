"".foo STEXT size=126 args=0x8 locals=0x28
	0x0000 00000 (main.go:3)	TEXT	"".foo(SB), $40-8 //40 是栈帧大小（局部变量+可能需要的额外调用函数的参数空间的总大小），8是参数及返回值大小(可选的)
	0x0000 00000 (main.go:3)	MOVQ	(TLS), CX //调度相关指令
	0x0009 00009 (main.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:3)	JLS	119
	0x000f 00015 (main.go:3)	SUBQ	$40, SP //将SP栈顶指针下移40字节，对SP做减法，即分配函数栈帧
	0x0013 00019 (main.go:3)	MOVQ	BP, 32(SP)//将BP寄存器的值入栈
	0x0018 00024 (main.go:3)	LEAQ	32(SP), BP//将新的栈顶地址保存到BP寄存器
	0x001d 00029 (main.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB) //gc 相关
	0x001d 00029 (main.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:3)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (main.go:3)	PCDATA	$2, $0 //gc 相关
	0x001d 00029 (main.go:3)	PCDATA	$0, $0
	0x001d 00029 (main.go:3)	MOVQ	$0, "".r+48(SP)//r=0
	0x0026 00038 (main.go:4)	MOVQ	$5, "".t+24(SP)//t=5
	0x002f 00047 (main.go:7)	PCDATA	$2, $1
	0x002f 00047 (main.go:7)	LEAQ	"".t+24(SP), AX//获取t的地址并搬移到AX
	0x0034 00052 (main.go:7)	PCDATA	$2, $0
	0x0034 00052 (main.go:7)	MOVQ	AX, 16(SP)//将AX寄存器的值搬移到栈顶16位移处
	0x0039 00057 (main.go:5)	MOVL	$8, (SP)
	0x0040 00064 (main.go:5)	PCDATA	$2, $1
	0x0040 00064 (main.go:5)	LEAQ	"".foo.func1·f(SB), AX
	0x0047 00071 (main.go:5)	PCDATA	$2, $0
	0x0047 00071 (main.go:5)	MOVQ	AX, 8(SP)
	0x004c 00076 (main.go:5)	CALL	runtime.deferproc(SB)
	0x0051 00081 (main.go:5)	TESTL	AX, AX
	0x0053 00083 (main.go:5)	JNE	103
	0x0055 00085 (main.go:5)	JMP	87
	0x0057 00087 (main.go:8)	XCHGL	AX, AX
	0x0058 00088 (main.go:8)	CALL	runtime.deferreturn(SB)
	0x005d 00093 (main.go:8)	MOVQ	32(SP), BP
	0x0062 00098 (main.go:8)	ADDQ	$40, SP
	0x0066 00102 (main.go:8)	RET
	0x0067 00103 (main.go:5)	XCHGL	AX, AX
	0x0068 00104 (main.go:5)	CALL	runtime.deferreturn(SB)
	0x006d 00109 (main.go:5)	MOVQ	32(SP), BP
	0x0072 00114 (main.go:5)	ADDQ	$40, SP //对sp做加法，清除函数栈帧
	0x0076 00118 (main.go:5)	RET
	0x0077 00119 (main.go:5)	NOP
	0x0077 00119 (main.go:3)	PCDATA	$0, $-1
	0x0077 00119 (main.go:3)	PCDATA	$2, $-1
	0x0077 00119 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x007c 00124 (main.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 68 48  eH..%....H;a.vhH
	0x0010 83 ec 28 48 89 6c 24 20 48 8d 6c 24 20 48 c7 44  ..(H.l$ H.l$ H.D
	0x0020 24 30 00 00 00 00 48 c7 44 24 18 05 00 00 00 48  $0....H.D$.....H
	0x0030 8d 44 24 18 48 89 44 24 10 c7 04 24 08 00 00 00  .D$.H.D$...$....
	0x0040 48 8d 05 00 00 00 00 48 89 44 24 08 e8 00 00 00  H......H.D$.....
	0x0050 00 85 c0 75 12 eb 00 90 e8 00 00 00 00 48 8b 6c  ...u.........H.l
	0x0060 24 20 48 83 c4 28 c3 90 e8 00 00 00 00 48 8b 6c  $ H..(.......H.l
	0x0070 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb 82        $ H..(........
	rel 5+4 t=16 TLS+0
	rel 67+4 t=15 "".foo.func1·f+0
	rel 77+4 t=8 runtime.deferproc+0
	rel 89+4 t=8 runtime.deferreturn+0
	rel 105+4 t=8 runtime.deferreturn+0
	rel 120+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=89 args=0x0 locals=0x18
	0x0000 00000 (main.go:11)	TEXT	"".main(SB), $24-0
	0x0000 00000 (main.go:11)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:11)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:11)	JLS	82
	0x000f 00015 (main.go:11)	SUBQ	$24, SP
	0x0013 00019 (main.go:11)	MOVQ	BP, 16(SP)
	0x0018 00024 (main.go:11)	LEAQ	16(SP), BP
	0x001d 00029 (main.go:11)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:11)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (main.go:12)	PCDATA	$2, $0
	0x001d 00029 (main.go:12)	PCDATA	$0, $0
	0x001d 00029 (main.go:12)	CALL	"".foo(SB)
	0x0022 00034 (main.go:12)	MOVQ	(SP), AX
	0x0026 00038 (main.go:12)	MOVQ	AX, ""..autotmp_0+8(SP)
	0x002b 00043 (main.go:12)	CALL	runtime.printlock(SB)
	0x0030 00048 (main.go:12)	MOVQ	""..autotmp_0+8(SP), AX
	0x0035 00053 (main.go:12)	MOVQ	AX, (SP)
	0x0039 00057 (main.go:12)	CALL	runtime.printint(SB)
	0x003e 00062 (main.go:12)	CALL	runtime.printnl(SB)
	0x0043 00067 (main.go:12)	CALL	runtime.printunlock(SB)
	0x0048 00072 (main.go:13)	MOVQ	16(SP), BP
	0x004d 00077 (main.go:13)	ADDQ	$24, SP
	0x0051 00081 (main.go:13)	RET
	0x0052 00082 (main.go:13)	NOP
	0x0052 00082 (main.go:11)	PCDATA	$0, $-1
	0x0052 00082 (main.go:11)	PCDATA	$2, $-1
	0x0052 00082 (main.go:11)	CALL	runtime.morestack_noctxt(SB)
	0x0057 00087 (main.go:11)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 43 48  eH..%....H;a.vCH
	0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 e8 00 00  ...H.l$.H.l$....
	0x0020 00 00 48 8b 04 24 48 89 44 24 08 e8 00 00 00 00  ..H..$H.D$......
	0x0030 48 8b 44 24 08 48 89 04 24 e8 00 00 00 00 e8 00  H.D$.H..$.......
	0x0040 00 00 00 e8 00 00 00 00 48 8b 6c 24 10 48 83 c4  ........H.l$.H..
	0x0050 18 c3 e8 00 00 00 00 eb a7                       .........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=8 "".foo+0
	rel 44+4 t=8 runtime.printlock+0
	rel 58+4 t=8 runtime.printint+0
	rel 63+4 t=8 runtime.printnl+0
	rel 68+4 t=8 runtime.printunlock+0
	rel 83+4 t=8 runtime.morestack_noctxt+0
"".foo.func1 STEXT nosplit size=21 args=0x8 locals=0x0
	0x0000 00000 (main.go:5)	TEXT	"".foo.func1(SB), NOSPLIT, $0-8
	0x0000 00000 (main.go:5)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0000 00000 (main.go:5)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0000 00000 (main.go:5)	FUNCDATA	$3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x0000 00000 (main.go:6)	PCDATA	$2, $1
	0x0000 00000 (main.go:6)	PCDATA	$0, $0
	0x0000 00000 (main.go:6)	MOVQ	"".&t+8(SP), AX
	0x0005 00005 (main.go:6)	PCDATA	$2, $2
	0x0005 00005 (main.go:6)	PCDATA	$0, $1
	0x0005 00005 (main.go:6)	MOVQ	"".&t+8(SP), CX
	0x000a 00010 (main.go:6)	PCDATA	$2, $1
	0x000a 00010 (main.go:6)	MOVQ	(CX), CX
	0x000d 00013 (main.go:6)	ADDQ	$5, CX
	0x0011 00017 (main.go:6)	PCDATA	$2, $0
	0x0011 00017 (main.go:6)	MOVQ	CX, (AX)
	0x0014 00020 (main.go:7)	RET
	0x0000 48 8b 44 24 08 48 8b 4c 24 08 48 8b 09 48 83 c1  H.D$.H.L$.H..H..
	0x0010 05 48 89 08 c3                                   .H...
"".init STEXT size=95 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	88
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	CMPB	"".initdone·(SB), $1
	0x0022 00034 (<autogenerated>:1)	JHI	38
	0x0024 00036 (<autogenerated>:1)	JMP	47
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $0
	0x002f 00047 (<autogenerated>:1)	PCDATA	$0, $0
	0x002f 00047 (<autogenerated>:1)	CMPB	"".initdone·(SB), $1
	0x0036 00054 (<autogenerated>:1)	JEQ	58
	0x0038 00056 (<autogenerated>:1)	JMP	65
	0x003a 00058 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x003f 00063 (<autogenerated>:1)	UNDEF
	0x0041 00065 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0048 00072 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x004f 00079 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0053 00083 (<autogenerated>:1)	ADDQ	$8, SP
	0x0057 00087 (<autogenerated>:1)	RET
	0x0058 00088 (<autogenerated>:1)	NOP
	0x0058 00088 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0058 00088 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0058 00088 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x005d 00093 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 49 48  eH..%....H;a.vIH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 80 3d 00 00 00  ...H.,$H.,$.=...
	0x0020 00 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08 c3 80  ..w...H.,$H.....
	0x0030 3d 00 00 00 00 01 74 02 eb 07 e8 00 00 00 00 0f  =.....t.........
	0x0040 0b c6 05 00 00 00 00 01 c6 05 00 00 00 00 02 48  ...............H
	0x0050 8b 2c 24 48 83 c4 08 c3 e8 00 00 00 00 eb a1     .,$H...........
	rel 5+4 t=16 TLS+0
	rel 29+4 t=15 "".initdone·+-1
	rel 49+4 t=15 "".initdone·+-1
	rel 59+4 t=8 runtime.throwinit+0
	rel 67+4 t=15 "".initdone·+-1
	rel 74+4 t=15 "".initdone·+-1
	rel 89+4 t=8 runtime.morestack_noctxt+0
go.loc."".foo SDWARFLOC size=0
go.info."".foo SDWARFINFO size=54
	0x0000 02 22 22 2e 66 6f 6f 00 00 00 00 00 00 00 00 00  ."".foo.........
	0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0e  ................
	0x0020 72 00 01 03 00 00 00 00 01 9c 09 74 00 04 00 00  r..........t....
	0x0030 00 00 02 91 68 00                                ....h.
	rel 8+8 t=1 "".foo+0
	rel 16+8 t=1 "".foo+126
	rel 26+4 t=29 gofile../data/app/go/src/gosrc/practice/func/defer/asm/main.go+0
	rel 36+4 t=28 go.info.int+0
	rel 46+4 t=28 go.info.int+0
go.range."".foo SDWARFRANGE size=0
go.isstmt."".foo SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 17 01 05 02 07 01 11 02 02 01 04  ................
	0x0010 02 06 01 10 02 11 00                             .......
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=33
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+89
	rel 27+4 t=29 gofile../data/app/go/src/gosrc/practice/func/defer/asm/main.go+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 05 01 26 02 11 00                 .......&...
go.loc."".foo.func1 SDWARFLOC size=0
go.info."".foo.func1 SDWARFINFO size=50
	0x0000 02 22 22 2e 66 6f 6f 2e 66 75 6e 63 31 00 00 00  ."".foo.func1...
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 9c  ................
	0x0020 00 00 00 00 01 0e 26 74 00 00 05 00 00 00 00 01  ......&t........
	0x0030 9c 00                                            ..
	rel 14+8 t=1 "".foo.func1+0
	rel 22+8 t=1 "".foo.func1+21
	rel 32+4 t=29 gofile../data/app/go/src/gosrc/practice/func/defer/asm/main.go+0
	rel 43+4 t=28 go.info.*int+0
go.range."".foo.func1 SDWARFRANGE size=0
go.isstmt."".foo.func1 SDWARFMISC size=0
	0x0000 04 05 01 0f 02 01 00                             .......
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+95
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 04 02 09 01 10 02 09 01 10  ................
	0x0010 02 07 00                                         ...
"".initdone· SNOPTRBSS size=1
"".foo.func1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".foo.func1+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
