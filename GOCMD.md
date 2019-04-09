

// $GODEV/bin/go 使用开发模式
// Gosched 让出cpu
// Goexit 退出当前goroutine(但是defer语句会照常执行)
// GOMAXPROCS  设置逻辑处理器数量，即最大的可同时使用的CPU核数 set env: export GOMAXPROCS=1
// NumCPU 返回当前系统的CPU核数量

##  build
`go build -o main -gcflags "all=-N -l"`

### 打印编译信息
  `go tool compile -d=slice,append,gcprog,closure,export,wb -o t.a main.go`

## print asm

`go tool objdump -S -s "main\.main" main`

` go tool compile -N -l -S main.go > main.s`

`go build -gcflags -S main.go  more asm`

## gctrace
`go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main `
`go build -o main /data/app/go/src/gosrc/demo/runtime/goroutine/for/main.go && GODEBUG=gctrace=1,gccheckmark=0  ./main`

##  goroutine schedtrace
`go build -o main demo/runtime/goroutine/main.go && GODEBUG=schedtrace=10000,scheddetail=1 ./main`


`sudo go tool dist test -v -run=^go_test:runtime$`

## test
``go tool compile -N -l -S demo/type/interface/struct/main.go > main.s`

`go build -o main -gcflags "all=-N -l" demo/type/interface/struct/main.go && go tool objdump -S -s "main\.main" main > main.s`

### list release tag
`go list  -f '{{context.ReleaseTags}}' runtime | grep go1.10`
 [go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11]

`go build -o main -gcflags "all=-N -l -m" demo/runtime/memory/closure/main.go`

### Visualise Go program GC trace data in real time 
https://github.com/davecheney/gcvis

`allocfreetrace: 设置 allocfreetrace=1 会监控每次分配，但因每次分配和释放的栈信息（stack trace）`

`cgocheck: 设置 cgocheck=0 禁用所有cgo检查将Go指针传递给非Go代码是否正确。cgocheck=1 (缺省值) 轻量级检查。cgocheck=2 重量级检查。`

`efence: 设置 efence=1 导致分配器 allocator将每个对象分配在一个唯一的页page上，地址不重用。`

`gccheckmark: 设置 gccheckmark=1 允许垃圾回收器执行并发mark阶段的校验。会导致Stop The World。`

`gcpacertrace: 设置 gcpacertrace=1 会让来几回收器打印出concurrent pacer的内部状态。`

`gcshrinkstackoff: 设置 gcshrinkstackoff=1 则禁止将 goroutines 的栈缩小为更小栈。`

`gcstackbarrieroff: 设置 gcstackbarrieroff=1 禁用stack barriers，会影响垃圾回收器的重复搜索栈的功能。`

`gcstackbarrierall: 设置 gcstackbarrierall=1 会为每个栈帧安装一 stack barriers。`

`gcstoptheworld: 设置 gcstoptheworld=1 则禁用并发垃圾回收,每次回收都会触发STW。设置gcstoptheworld=2则禁用垃圾回收后的concurrent sweeping。`

`gctrace: 设置 gctrace=1导致每次垃圾回收器触发一行日志，包含内存回收的概要信息和暂停的时间。设置gctrace=2起同样的效果，but also repeats each collection。格式如下：`

    gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
where the fields are as follows:
    gc #        GC id,每次GC加一
    @#s         程序启动后的时间，单位秒
    #%          程序启动后GC所用的时间比
    #+...+#     此次GC所用的wall-clock/CPU时间
    #->#-># MB  GC开始时的堆大小, GC结束时的堆大小, 活着的(live)堆大小
    # MB goal   总的堆大小
    # P         CPU使用数
垃圾回收分为下面的几个阶段：stop-the-world (STW) sweep termination, concurrent
mark and scan, and STW mark termination。 mark/scan的CPU时间又分为 assist time (GC performed in
line with allocation), background GC time, and idle GC time。
如果日志后以"(forced)"结尾,则GC通过runtime.GC()调用执行，此时所有的阶段都是STW.

`memprofilerate: 设置 memprofilerate=X 会更新runtime.MemProfileRate的值。0则禁用这个profie。`

`invalidptr: 默认设为invalidptr=1, 如果指针被赋予一个无效值,会引起程序的崩溃，设置该值为0，会停止该检查，0只能临时用于查找bug，真正的解决方法是不要把整数类型的值存在指针变量里面。`

`sbrk: 设置 sbrk=1 会使用实验性的实现替换memory allocator 和 garbage collector。`

`scavenge: scavenge=1 允许heap scavenger的debug模式。`

`scheddetail: 设置 schedtrace=X 和 scheddetail=1 会导致goroutine调度器每个X毫秒输出多行调度信息。`

`schedtrace: 设置 schedtrace=X导致调度器每个X秒输出一行调度器的概要信息。`

```