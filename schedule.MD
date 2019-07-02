### sysmon
```text sysmon是一个由runtime启动的M，也叫监控线程，它无需P也可以运行，它每20us~10ms唤醒一次，主要执行:
释放闲置超过5分钟的span物理内存；
如果超过2分钟没有垃圾回收，强制执行；
将长时间未处理的netpoll结果添加到任务队列；
向长时间运行的G任务发出抢占调度；
收回因syscall长时间阻塞的P；
```

追踪：
all goroutines are asleep - deadlock! 
runtime.checkdead() ./go/src/runtime/proc.go:4225 (PC: 0x10370d3)

0  0x0000000001038949 in runtime.mput
   at ./go/src/runtime/proc.go:4654
1  0x00000000010314b8 in runtime.stopm
   at ./go/src/runtime/proc.go:2015
2  0x000000000103265c in runtime.findrunnable
   at ./go/src/runtime/proc.go:2488
3  0x000000000103316a in runtime.schedule
   at ./go/src/runtime/proc.go:2614
4  0x000000000103348e in runtime.park_m
   at ./go/src/runtime/proc.go:2677
5  0x000000000105594b in runtime.mcall
   at ./go/src/runtime/asm_amd64.s:299

### 参考文章

- [Go 调度模型](https://wudaijun.com/2018/01/go-scheduler)
- [深入Golang调度器之GMP模型](https://www.jianshu.com/p/abe79d86ff27)
- [go语言调度器源代码情景分析](https://mp.weixin.qq.com/mp/homepage?__biz=MzU1OTg5NDkzOA==&hid=1&sn=8fc2b63f53559bc0cee292ce629c4788&scene=25#wechat_redirect)

- [Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs)
- [The Go scheduler](http://morsmachine.dk/go-scheduler)
- [也谈goroutine调度器](http://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/)
- [Scalable Go Scheduler Design Doc](https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit#)
- [Go Preemptive Scheduler Design Doc](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#!)
- [goroutine背后的系统知识](http://www.sizeofvoid.net/goroutine-under-the-hood/)
- [NUMA-aware scheduler for Go](https://docs.google.com/document/u/0/d/1d3iI2QWURgDIsSR6G2275vMeQ_X7w-qxM2Vp7iGwwuM/pub)
- [The Go netpoller](http://morsmachine.dk/netpoller)
- [Go runtime doc](https://golang.org/pkg/runtime/)
- [Work Stealing论文](http://supertech.csail.mit.edu/papers/steal.pdf)
- [Work Stealing维基百科](https://en.wikipedia.org/wiki/Work_stealing)






