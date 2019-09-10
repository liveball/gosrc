
### Quick intro to go assembly
https://blog.hackercat.ninja/post/quick_intro_to_go_assembly/

### Go 语言汇编快速入门
https://studygolang.com/articles/12828?fr=sidebar

### plan9  汇编
https://9p.io/sys/doc/asm.pdf


#编译汇编

nasm -f elf64 -o hello.o hello.asm

nasm -f elf32 -o hello.o hello.asm

ld -o hello hello.o

### 编译
`nasm -g -F dwarf -f elf64 -o hello.o hello.s`

### 链接
`ld -o hel hello.o`

### nm命令主要是用来列出某些文件中的符号（说白了就是一些函数和全局变量等）
`nm hel`

### 删除标签
`strip hel`

### 反汇编

`objdump -d -M intel hel`

```asm hel：     文件格式 elf64-x86-64


Disassembly of section .text:

00000000004000b0 <main>:
  4000b0:	b8 01 00 00 00       	mov    eax,0x1
  4000b5:	bf 01 00 00 00       	mov    edi,0x1
  4000ba:	48 be dc 00 60 00 00 	movabs rsi,0x6000dc
  4000c1:	00 00 00 
  4000c4:	ba 0e 00 00 00       	mov    edx,0xe
  4000c9:	0f 05                	syscall 
  4000cb:	eb 00                	jmp    4000cd <main.exit>

00000000004000cd <main.exit>:
  4000cd:	b8 3c 00 00 00       	mov    eax,0x3c
  4000d2:	48 31 ff             	xor    rdi,rdi
  4000d5:	0f 05                	syscall 

00000000004000d7 <_start>:
  4000d7:	eb d7                	jmp    4000b0 <main>
  
```

`readelf -s hello.o`

Symbol table '.symtab' contains 11 entries:
   Num:    Value          Size Type    Bind   Vis      Ndx Name
     0: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT  UND 
     1: 0000000000000000     0 FILE    LOCAL  DEFAULT  ABS hello.s
     2: 0000000000000000     0 SECTION LOCAL  DEFAULT    1 
     3: 0000000000000000     0 SECTION LOCAL  DEFAULT    2 
     4: 0000000000000000     1 OBJECT  LOCAL  DEFAULT    1 hello
     5: 0000000000000000     0 NOTYPE  LOCAL  DEFAULT    2 main
     6: 000000000000001d     0 NOTYPE  LOCAL  DEFAULT    2 main.exit
     7: 0000000000000000     0 SECTION LOCAL  DEFAULT   10 
     8: 0000000000000000     0 SECTION LOCAL  DEFAULT   12 
     9: 0000000000000000     0 SECTION LOCAL  DEFAULT   13 
    10: 0000000000000027     0 NOTYPE  GLOBAL DEFAULT    2 _start
    
`readelf -l hel`

Elf 文件类型为 EXEC (可执行文件)
Entry point 0x4000d7
There are 2 program headers, starting at offset 64

程序头：
  Type           Offset             VirtAddr           PhysAddr
                 FileSiz            MemSiz              Flags  Align
  LOAD           0x0000000000000000 0x0000000000400000 0x0000000000400000
                 0x00000000000000d9 0x00000000000000d9  R E    0x200000
  LOAD           0x00000000000000dc 0x00000000006000dc 0x00000000006000dc
                 0x000000000000000d 0x000000000000000d  RW     0x200000

 Section to Segment mapping:
  段节...
   00     .text 
   01     .data 
