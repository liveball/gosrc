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


// func setg(gg *g)
// set g. for use by needm.
TEXT ·setg(SB), NOSPLIT, $0-8
	MOVQ	gg+0(FP), BX
#ifdef GOOS_windows
	CMPQ	BX, $0
	JNE	settls
	MOVQ	$0, 0x28(GS)
	RET
settls:
	MOVQ	g_m(BX), AX
	LEAQ	m_tls(AX), AX
	MOVQ	AX, 0x28(GS)
#endif
	get_tls(CX)
	MOVQ	BX, g(CX)
	RET

// func StringLen(s string) int
TEXT ·StringLen(SB), NOSPLIT, $0-24
	MOVQ s_len+8(FP), AX
	MOVQ AX, ret+16(FP)
	RET

// func String(s string) string
TEXT ·String(SB), NOSPLIT, $0-32
    MOVQ s_str+0(FP), AX
    MOVQ s_len+8(FP), BX
   	MOVQ AX, ret+16(FP)
   	MOVQ BX, ret+24(FP)
    RET

//func NewStudent(s Student) Student
TEXT ·NewStudent(SB), NOSPLIT, $0-48
	MOVQ  s_Name_str+0(FP), AX
	MOVQ  s_Name_len+8(FP), BX
	MOVQ  s_Age+16(FP), CX

	MOVQ  AX, ret_Name_str+24(FP)
	MOVQ BX, ret_Name_len+32(FP)
	MOVQ CX, ret_Age+40(FP)
	RET

//func NewStudentPtr(s *Student) *Student
TEXT ·NewStudentPtr(SB), NOSPLIT, $0-16
    MOVQ  s+0(FP), AX
	MOVQ  AX, ret+8(FP)
	RET

//地址运算也是用 lea 指令，英文原意为Load Effective Address，amd64 平台地址都是8个字节，所以直接就用LEAQ就好：
//MOVQ 在寄存器加偏移的情况下MOVQ会对地址进行解引用

//func UpStudentPtr(s *Student, name string, age int)
TEXT ·UpStudentPtr(SB), NOSPLIT, $0-0
    //BX 存放s指针
    MOVQ  s+0(FP), BX    // FP+0  为参数s，将其值拷贝到寄存器BX中

    //AX 存name CX 存age
    MOVQ  name+8(FP), AX // FP+8  为参数name，将其值拷贝到寄存器AX中
    MOVQ  name+24(FP), CX // FP+24  为参数age，将其值拷贝到寄存器CX中

    // 先用DX存name的地址，然后把DX放入s内存前0-16 bytes
    LEAQ   (AX), DX //DX = &AX 将AX的地址，放入DX
    MOVQ   DX, (BX) //*BX = DX 将DX寄存器中的值， 放入BX指向的内存区域16byte

    // 先用DX存age的地址，然后把DX放入s内存前16-24bytes
    LEAQ   (CX), DX //DX = &AX 将AX的地址，放入DX
    MOVQ   DX, 16(BX) //*(BX+16) = DX 将DX寄存器中的值， 放入BX指向的内存区域8byte
	RET


// func StudentName(s *Student) string
TEXT ·StudentName(SB), NOSPLIT, $0-16
	MOVQ	s+0(FP), BX
	MOVQ	s_Name(BX), AX
	RET

