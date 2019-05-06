### 打印所有ssa生成阶段代码，并且生成交互式ssa.html
`GOSSAFUNC=main GOOS=linux GOARCH=amd64 go build -gcflags -S main.go` 

ref:
https://www.yiqishare.com/news/13.html
https://toutiao.io/posts/572721/app_preview

###汇编
参考：
A Manual for the Plan 9 assembler --Rob Pike
https://9p.io/sys/doc/asm.html

Go & Assembly
https://www.doxsey.net/blog/go-and-assembly

《plan9 汇编入门，带你打通应用和底层》
https://github.com/developer-learning/reading-go/issues/186

golang汇编基础知识
http://guidao.github.io/asm.html

汇编中movl,movw,movb的作用
https://blog.csdn.net/m0_37806112/article/details/80549927

汇编角度看 Slice，一个新的世界
https://mp.weixin.qq.com/s/MHxPZS3MVRG0DjjKAyWLmg

