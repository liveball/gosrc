#include "textflag.h"

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
TEXT ·NewStudent(SB), NOSPLIT, $0-64
	MOVQ  s_Name_str+0(FP), AX
	MOVQ  s_Name_len+8(FP), BX
	MOVQ  s_Age+16(FP), CX
	MOVQ  s_Card+24(FP), DX

	MOVQ  AX, ret_Name_str+32(FP)
	MOVQ  BX, ret_Name_len+40(FP)
	MOVQ  CX, ret_Age+48(FP)
	MOVQ  DX, ret_Card+56(FP)
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


// func StudentPtr(s *Student) unsafe.Pointer
TEXT ·StudentPtr(SB), NOSPLIT, $0-16
	MOVQ	s+0(FP), BX
   	MOVQ    BX, ret+8(FP)
	RET

// func StudentAge(s *Student) int
TEXT ·StudentAge(SB), NOSPLIT, $0-16
     MOVQ s+0(FP), BX

     MOVQ 16(BX), AX

     MOVQ  AX, ret+8(FP)

     RET

// func StudentName(s *Student) string
TEXT ·StudentName(SB), NOSPLIT, $0-24
     MOVQ s+0(FP), BX

     MOVQ (BX), AX  // string 第一个字段地址
     MOVQ 8(BX), CX // string 第二个字段地址

     MOVQ AX, ret+8(FP)  //返回到栈帧 8字节处
     MOVQ CX, ret+16(FP) //返回到栈帧 16字节处

     RET

