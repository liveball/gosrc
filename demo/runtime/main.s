"".main STEXT size=193 args=0x0 locals=0x28
	0x0000 00000 (main.go:8)	TEXT	"".main(SB), $40-0
	0x0000 00000 (main.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:8)	JLS	183
	0x0013 00019 (main.go:8)	SUBQ	$40, SP
	0x0017 00023 (main.go:8)	MOVQ	BP, 32(SP)
	0x001c 00028 (main.go:8)	LEAQ	32(SP), BP
	0x0021 00033 (main.go:8)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (main.go:8)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (main.go:8)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0021 00033 (main.go:9)	PCDATA	$2, $0
	0x0021 00033 (main.go:9)	PCDATA	$0, $0
	0x0021 00033 (main.go:9)	CALL	runtime.printlock(SB)
	0x0026 00038 (main.go:9)	MOVQ	$1, (SP)
	0x002e 00046 (main.go:9)	CALL	runtime.printint(SB)
	0x0033 00051 (main.go:9)	CALL	runtime.printnl(SB)
	0x0038 00056 (main.go:9)	CALL	runtime.printunlock(SB)
	0x003d 00061 (main.go:11)	PCDATA	$2, $1
	0x003d 00061 (main.go:11)	LEAQ	type.sync.WaitGroup(SB), AX
	0x0044 00068 (main.go:11)	PCDATA	$2, $0
	0x0044 00068 (main.go:11)	MOVQ	AX, (SP)
	0x0048 00072 (main.go:11)	CALL	runtime.newobject(SB)
	0x004d 00077 (main.go:11)	PCDATA	$2, $1
	0x004d 00077 (main.go:11)	MOVQ	8(SP), AX
	0x0052 00082 (main.go:11)	PCDATA	$0, $1
	0x0052 00082 (main.go:11)	MOVQ	AX, "".&wg+24(SP)
	0x0057 00087 (main.go:11)	MOVQ	$0, (AX)
	0x005e 00094 (main.go:11)	PCDATA	$2, $0
	0x005e 00094 (main.go:11)	MOVQ	$0, 4(AX)
	0x0066 00102 (main.go:12)	PCDATA	$2, $1
	0x0066 00102 (main.go:12)	MOVQ	"".&wg+24(SP), AX
	0x006b 00107 (main.go:12)	PCDATA	$2, $0
	0x006b 00107 (main.go:12)	MOVQ	AX, (SP)
	0x006f 00111 (main.go:12)	MOVQ	$1, 8(SP)
	0x0078 00120 (main.go:12)	CALL	sync.(*WaitGroup).Add(SB)
	0x007d 00125 (main.go:13)	PCDATA	$2, $1
	0x007d 00125 (main.go:13)	MOVQ	"".&wg+24(SP), AX
	0x0082 00130 (main.go:16)	PCDATA	$2, $0
	0x0082 00130 (main.go:16)	MOVQ	AX, 16(SP)
	0x0087 00135 (main.go:13)	MOVL	$8, (SP)
	0x008e 00142 (main.go:13)	PCDATA	$2, $1
	0x008e 00142 (main.go:13)	LEAQ	"".main.func1·f(SB), AX
	0x0095 00149 (main.go:13)	PCDATA	$2, $0
	0x0095 00149 (main.go:13)	MOVQ	AX, 8(SP)
	0x009a 00154 (main.go:13)	CALL	runtime.newproc(SB)
	0x009f 00159 (main.go:18)	PCDATA	$2, $1
	0x009f 00159 (main.go:18)	PCDATA	$0, $0
	0x009f 00159 (main.go:18)	MOVQ	"".&wg+24(SP), AX
	0x00a4 00164 (main.go:18)	PCDATA	$2, $0
	0x00a4 00164 (main.go:18)	MOVQ	AX, (SP)
	0x00a8 00168 (main.go:18)	CALL	sync.(*WaitGroup).Wait(SB)
	0x00ad 00173 (main.go:19)	MOVQ	32(SP), BP
	0x00b2 00178 (main.go:19)	ADDQ	$40, SP
	0x00b6 00182 (main.go:19)	RET
	0x00b7 00183 (main.go:19)	NOP
	0x00b7 00183 (main.go:8)	PCDATA	$0, $-1
	0x00b7 00183 (main.go:8)	PCDATA	$2, $-1
	0x00b7 00183 (main.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x00bc 00188 (main.go:8)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a4  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
	0x0020 20 e8 00 00 00 00 48 c7 04 24 01 00 00 00 e8 00   .....H..$......
	0x0030 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8d 05  .............H..
	0x0040 00 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b 44  ....H..$.....H.D
	0x0050 24 08 48 89 44 24 18 48 c7 00 00 00 00 00 48 c7  $.H.D$.H......H.
	0x0060 40 04 00 00 00 00 48 8b 44 24 18 48 89 04 24 48  @.....H.D$.H..$H
	0x0070 c7 44 24 08 01 00 00 00 e8 00 00 00 00 48 8b 44  .D$..........H.D
	0x0080 24 18 48 89 44 24 10 c7 04 24 08 00 00 00 48 8d  $.H.D$...$....H.
	0x0090 05 00 00 00 00 48 89 44 24 08 e8 00 00 00 00 48  .....H.D$......H
	0x00a0 8b 44 24 18 48 89 04 24 e8 00 00 00 00 48 8b 6c  .D$.H..$.....H.l
	0x00b0 24 20 48 83 c4 28 c3 e8 00 00 00 00 e9 3f ff ff  $ H..(.......?..
	0x00c0 ff                                               .
	rel 5+4 t=16 TLS+0
	rel 34+4 t=8 runtime.printlock+0
	rel 47+4 t=8 runtime.printint+0
	rel 52+4 t=8 runtime.printnl+0
	rel 57+4 t=8 runtime.printunlock+0
	rel 64+4 t=15 type.sync.WaitGroup+0
	rel 73+4 t=8 runtime.newobject+0
	rel 121+4 t=8 sync.(*WaitGroup).Add+0
	rel 145+4 t=15 "".main.func1·f+0
	rel 155+4 t=8 runtime.newproc+0
	rel 169+4 t=8 sync.(*WaitGroup).Wait+0
	rel 184+4 t=8 runtime.morestack_noctxt+0
"".main.func1 STEXT size=88 args=0x8 locals=0x10
	0x0000 00000 (main.go:13)	TEXT	"".main.func1(SB), $16-8
	0x0000 00000 (main.go:13)	MOVQ	(TLS), CX
	0x0009 00009 (main.go:13)	CMPQ	SP, 16(CX)
	0x000d 00013 (main.go:13)	JLS	81
	0x000f 00015 (main.go:13)	SUBQ	$16, SP
	0x0013 00019 (main.go:13)	MOVQ	BP, 8(SP)
	0x0018 00024 (main.go:13)	LEAQ	8(SP), BP
	0x001d 00029 (main.go:13)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x001d 00029 (main.go:13)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x001d 00029 (main.go:13)	FUNCDATA	$3, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x001d 00029 (main.go:14)	PCDATA	$2, $0
	0x001d 00029 (main.go:14)	PCDATA	$0, $0
	0x001d 00029 (main.go:14)	CALL	runtime.printlock(SB)
	0x0022 00034 (main.go:14)	MOVQ	$2, (SP)
	0x002a 00042 (main.go:14)	CALL	runtime.printint(SB)
	0x002f 00047 (main.go:14)	CALL	runtime.printnl(SB)
	0x0034 00052 (main.go:14)	CALL	runtime.printunlock(SB)
	0x0039 00057 (main.go:15)	PCDATA	$2, $1
	0x0039 00057 (main.go:15)	PCDATA	$0, $1
	0x0039 00057 (main.go:15)	MOVQ	"".&wg+24(SP), AX
	0x003e 00062 (main.go:15)	PCDATA	$2, $0
	0x003e 00062 (main.go:15)	MOVQ	AX, (SP)
	0x0042 00066 (main.go:15)	CALL	sync.(*WaitGroup).Done(SB)
	0x0047 00071 (main.go:16)	MOVQ	8(SP), BP
	0x004c 00076 (main.go:16)	ADDQ	$16, SP
	0x0050 00080 (main.go:16)	RET
	0x0051 00081 (main.go:16)	NOP
	0x0051 00081 (main.go:13)	PCDATA	$0, $-1
	0x0051 00081 (main.go:13)	PCDATA	$2, $-1
	0x0051 00081 (main.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x0056 00086 (main.go:13)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 42 48  dH..%....H;a.vBH
	0x0010 83 ec 10 48 89 6c 24 08 48 8d 6c 24 08 e8 00 00  ...H.l$.H.l$....
	0x0020 00 00 48 c7 04 24 02 00 00 00 e8 00 00 00 00 e8  ..H..$..........
	0x0030 00 00 00 00 e8 00 00 00 00 48 8b 44 24 18 48 89  .........H.D$.H.
	0x0040 04 24 e8 00 00 00 00 48 8b 6c 24 08 48 83 c4 10  .$.....H.l$.H...
	0x0050 c3 e8 00 00 00 00 eb a8                          ........
	rel 5+4 t=16 TLS+0
	rel 30+4 t=8 runtime.printlock+0
	rel 43+4 t=8 runtime.printint+0
	rel 48+4 t=8 runtime.printnl+0
	rel 53+4 t=8 runtime.printunlock+0
	rel 67+4 t=8 sync.(*WaitGroup).Done+0
	rel 82+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=100 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	93
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
	0x0048 00072 (<autogenerated>:1)	CALL	sync.init(SB)
	0x004d 00077 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x0054 00084 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0058 00088 (<autogenerated>:1)	ADDQ	$8, SP
	0x005c 00092 (<autogenerated>:1)	RET
	0x005d 00093 (<autogenerated>:1)	NOP
	0x005d 00093 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005d 00093 (<autogenerated>:1)	PCDATA	$2, $-1
	0x005d 00093 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0062 00098 (<autogenerated>:1)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 4e 48  dH..%....H;a.vNH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 80 3d 00 00 00  ...H.,$H.,$.=...
	0x0020 00 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08 c3 80  ..w...H.,$H.....
	0x0030 3d 00 00 00 00 01 74 02 eb 07 e8 00 00 00 00 0f  =.....t.........
	0x0040 0b c6 05 00 00 00 00 01 e8 00 00 00 00 c6 05 00  ................
	0x0050 00 00 00 02 48 8b 2c 24 48 83 c4 08 c3 e8 00 00  ....H.,$H.......
	0x0060 00 00 eb 9c                                      ....
	rel 5+4 t=16 TLS+0
	rel 29+4 t=15 "".initdone·+-1
	rel 49+4 t=15 "".initdone·+-1
	rel 59+4 t=8 runtime.throwinit+0
	rel 67+4 t=15 "".initdone·+-1
	rel 73+4 t=8 sync.init+0
	rel 79+4 t=15 "".initdone·+-1
	rel 94+4 t=8 runtime.morestack_noctxt+0
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=46
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 09 26 77 67 00 0b 00 00 00 00 02 91 68 00        .&wg........h.
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+193
	rel 27+4 t=29 gofile../data/app/go/src/go1.11.1/demo/runtime/main.go+0
	rel 38+4 t=28 go.info.*sync.WaitGroup+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 13 04 0e 03 05 01 17 02 07 01 22 02 05 01 12  ..........."....
	0x0010 02 11 01 11 02 05 01 09 02 14 00                 ...........
go.loc."".main.func1 SDWARFLOC size=0
go.info."".main.func1 SDWARFINFO size=52
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 31 00 00  ."".main.func1..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 9c 00 00 00 00 01 0e 26 77 67 00 00 0d 00 00 00  .......&wg......
	0x0030 00 01 9c 00                                      ....
	rel 15+8 t=1 "".main.func1+0
	rel 23+8 t=1 "".main.func1+88
	rel 33+4 t=29 gofile../data/app/go/src/go1.11.1/demo/runtime/main.go+0
	rel 45+4 t=28 go.info.*sync.WaitGroup+0
go.range."".main.func1 SDWARFRANGE size=0
go.isstmt."".main.func1 SDWARFMISC size=0
	0x0000 04 0f 04 0e 03 05 01 17 02 05 01 09 02 11 00     ...............
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+100
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 04 02 09 01 10 02 09 01 15  ................
	0x0010 02 07 00                                         ...
"".initdone· SNOPTRBSS size=1
"".main.func1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func1+0
type..importpath.sync. SRODATA dupok size=7
	0x0000 00 00 04 73 79 6e 63                             ...sync
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
