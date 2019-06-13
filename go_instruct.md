## //go: 指令 

### go:linkname-[链接指令](gosrc/practice/compile/go_instruct/go_link)
//go:linkname localname importpath.name
该指令指示编译器使用 importpath.name 作为源代码中声明为 localname 的变量或函数的目标文件符号名称。
但是由于这个伪指令，可以破坏类型系统和包模块化。因此只有引用了 unsafe 包才可以使用
简单来讲，就是 importpath.name 是 localname 的符号别名，编译器实际上会调用 localname 。但前提是使用了 unsafe包才能使用