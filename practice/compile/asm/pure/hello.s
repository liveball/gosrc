global _start

section .data
    hello : db `hello,world!\n`

section .text
    main:
        mov     rax,1
        mov     rdi,1
        mov     rsi,hello
        mov     rdx,14
        syscall
        jmp     .exit

 .exit:  ;main.exit
     mov        rax,60
     xor        rdi,rdi
     syscall


 _start:
        jmp main