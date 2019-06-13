## //go: 指令 
[practice/compile/go_instruct](https://github.com/liveball/gosrc/tree/master/practice/compile/go_instruct)

### go:linkname
//go:linkname localname importpath.name

```text 
该指令指示编译器使用 importpath.name 作为源代码中声明为 localname 的变量或函数的目标文件符号名称。
但是由于这个伪指令，可以破坏类型系统和包模块化。因此只有引用了 unsafe 包才可以使用
简单来讲，就是 importpath.name 是 localname 的符号别名，编译器实际上会调用 localname 。
但前提是使用了 unsafe包才能使用。
```

### //go:noescape

```text 
该指令指定下一个有声明但没有主体（意味着实现有可能不是 Go）的函数，不允许编译器对其做逃逸分析
一般情况下，该指令用于内存分配优化。因为编译器默认会进行逃逸分析，会通过规则判定一个变量是分配到堆上还是栈上。
但凡事有意外，一些函数虽然逃逸分析其是存放到堆上。但是对于我们来说，它是特别的。
我们就可以使用 go:noescape 指令强制要求编译器将其分配到函数栈上。
```


### //go:nosplit

```text 
该指令指定文件中声明的下一个函数不得包含堆栈溢出检查。简单来讲，就是这个函数跳过堆栈溢出的检查。
```

### //go:noinline

```text 
该指令表示该函数禁止进行内联。
```


### //go:yeswritebarrierrec

```text 
该指令与 go:nowritebarrierrec 相对，在标注 go:nowritebarrierrec 指令的函数上，遇到写屏障会产生错误。
而当编译器遇到 go:yeswritebarrierrec 指令时将会停止。
```

### //go:norace

```text 
该指令表示禁止进行竞态检测。而另外一种常见的形式就是在启动时执行 go run -race，能够检测应用程序中是否存在双向的数据竞争。
```

### //go:notinheap

```text 
该指令常用于类型声明，它表示这个类型不允许从 GC 堆上进行申请内存。在运行时中常用其来做较低层次的内部结构，避免调度器和内存分配中的写屏障。
```