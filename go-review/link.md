## 优秀commit

### cmd/compile: fix defer/deferreturn
https://go-review.googlesource.com/c/go/+/20486/

确保我们在所有路径上进行任何刚刚返回的清理
功能，包括恢复时。 每个出口路径应包括
deferreturn（如果有任何延迟）然后退出
代码（例如，将堆转义返回值复制回堆栈）。

引入一个Defer SSA块类型，它有两个输出边 - 一个
通过边缘（延迟成功排队）和一个
立即返回（延迟有一个成功的recover（）调用和
正常执行应该在返回点恢复）。


### cmd/compile: insert scheduling checks on some loop backedges
https://go-review.googlesource.com/c/go/+/33093/

GOEXPERIMENT = preemptibleloops

这是一种有助于研究和诊断的工具
GC和其他延迟问题，现在是目标STW GC延迟
大约是100微秒或更短。

它会降低大多数应用程序的性能
插入检查的开销。
