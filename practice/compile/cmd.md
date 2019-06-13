
### 伪寄存器
FP: Frame pointer: arguments and locals.(指向当前栈帧)    帧指针 –参数和局部变量–
PC: Program counter: jumps and branches.(指向指令地址)    程序计数器 –跳转和分支–
SB: Static base pointer: global symbols.(指向全局符号表)  静态基址指针 –全局符号–
SP: Stack pointer: top of stack.(指向当前栈顶部)          栈指针 –栈的顶端–.

BP:

### MOV
`MOVX source, destination`
source 和 destinatino 的值可以是内存地址，存储在内存中的数据值，指令语句中定义的数据值，或者是寄存器。

1、B用于8位的 1字节值      MOVB  1 byte
2、W用于16位的2字节值      MOVW  2 bytes
3、D用于32位的4字节值      MOVD  4 bytes
4、Q用于64位的8字节值      MOVQ  8 bytes
5、L用于32位的4字节值      MOVL  4 bytes  //L 这里代表 Long，4 字节的值，使用MOVL将一个小常量移动到寄存器中,当常量为正且适合32位时,32位低位辅助信号.


### XCHG 
`XCHGL	AX, AX`
Exchange Register/Memory with Register

### TESTL
`TESTL	AX, AX`

###RET
pop pc

###CALL
push pc
jmp to callee addr