### runtime/debug

>1.强制进行垃圾回收
```text
FreeOSMemory强制进行一次垃圾收集，以释放尽量多的内存回操作系统。（即使没有调用，运行时环境也会在后台任务里逐渐将内存释放给系统）
```
>2.设置垃圾回收的目标百分比  
```text
SetGCPercent设定垃圾收集的目标百分比：当新申请的内存大小占前次垃圾收集剩余可用内存大小的比率达到设定值时，就会触发垃圾收集。SetGCPercent返回之前的设定。初始值设定为环境变量GOGC的值；如果没有设置该环境变量，初始值为100。percent参数如果是负数值，
会关闭垃圾收集
```
>3.设置被单个go协程调用栈可使用的内存最大值  
```text
fmt.Println(debug.SetMaxStack(1)) //查看到默认系统为1000 000 000 字节
默认的设置32位系统是250MB,64位为1GB
``` 
>4.设置go程序可以使用的最大操作系统线程数  
```text
我们把程序的组大可使用的线程（不是协程）数设置为1，如果程序试图超过这个限制,程序就会崩溃，初始设置为10000个线程
什么时候会创建新的线程呢?
现有的线程阻塞，cgo或者runtime.LockOSThread函数阻塞其他go协程
```

>5.设置程序请求运行是只触发panic,而不崩溃  
```text
SetPanicOnFault控制程序在不期望（非nil）的地址出错时的运行时行为。这些错误一般是因为运行时内存破坏的bug引起的，因此默认反应是使程序崩溃。使用内存映射的文件或进行内存的不安全操作的程序可能会在非nil的地址出现错误；SetPanicOnFault允许这些程序请求运行时只触发一个panic，而不是崩溃。SetPanicOnFault只用于当前的go程，我们发现指针为nil 发生了panic 但是我们进行了恢复,程序继续执行
``` 

>6.垃圾收集信息的写入stats中  
>7.将内存分配堆和其中对象的描述写入文件中  
>8.获取go协程调用栈踪迹  
>9.将堆栈踪迹打印到标准错误   

```go
// ReadGCStats reads statistics about garbage collection into stats.
// The number of entries in the pause history is system-dependent;
// stats.Pause slice will be reused if large enough, reallocated otherwise.
// ReadGCStats may use the full capacity of the stats.Pause slice.
// If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles
// summarizing the distribution of pause time. For example, if
// len(stats.PauseQuantiles) is 5, it will be filled with the minimum,
// 25%, 50%, 75%, and maximum pause times.
func ReadGCStats(stats *GCStats) {
     func readGCStats(*[]time.Duration)

// FreeOSMemory forces a garbage collection followed by an
// attempt to return as much memory to the operating system
// as possible. (Even if this is not called, the runtime gradually
// returns memory to the operating system in a background task.)
func FreeOSMemory() 
    func freeOSMemory()

// SetMaxStack sets the maximum amount of memory that
// can be used by a single goroutine stack.
// If any goroutine exceeds this limit while growing its stack,
// the program crashes.
// SetMaxStack returns the previous setting.
// The initial setting is 1 GB on 64-bit systems, 250 MB on 32-bit systems.
//
// SetMaxStack is useful mainly for limiting the damage done by
// goroutines that enter an infinite recursion. It only limits future
// stack growth.
func SetMaxStack(bytes int) int {
     func setMaxStack(int) int

// SetGCPercent sets the garbage collection target percentage:
// a collection is triggered when the ratio of freshly allocated data
// to live data remaining after the previous collection reaches this percentage.
// SetGCPercent returns the previous setting.
// The initial setting is the value of the GOGC environment variable
// at startup, or 100 if the variable is not set.
// A negative percentage disables garbage collection.
func SetGCPercent(percent int) int {
    func setGCPercent(int32) int32

// SetPanicOnFault controls the runtime's behavior when a program faults
// at an unexpected (non-nil) address. Such faults are typically caused by
// bugs such as runtime memory corruption, so the default response is to crash
// the program. Programs working with memory-mapped files or unsafe
// manipulation of memory may cause faults at non-nil addresses in less
// dramatic situations; SetPanicOnFault allows such programs to request
// that the runtime trigger only a panic, not a crash.
// SetPanicOnFault applies only to the current goroutine.
// It returns the previous setting.
func SetPanicOnFault(enabled bool) bool {
    func setPanicOnFault(bool) bool

// SetMaxThreads sets the maximum number of operating system
// threads that the Go program can use. If it attempts to use more than
// this many, the program crashes.
// SetMaxThreads returns the previous setting.
// The initial setting is 10,000 threads.
//
// The limit controls the number of operating system threads, not the number
// of goroutines. A Go program creates a new thread only when a goroutine
// is ready to run but all the existing threads are blocked in system calls, cgo calls,
// or are locked to other goroutines due to use of runtime.LockOSThread.
//
// SetMaxThreads is useful mainly for limiting the damage done by
// programs that create an unbounded number of threads. The idea is
// to take down the program before it takes down the operating system.
func SetMaxThreads(threads int) int {
    func setMaxThreads(int) int
```

#参考：https://www.jianshu.com/p/0b3d11f7af57
