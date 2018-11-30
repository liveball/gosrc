### 编译main程序

http://ribrdb.github.io/lldb/

```go
go build -gcflags "-N -l" -o main main.go //关闭编译优化

go test -gcflags "-N -l" -c  -o lldb.test //注意test的包名必须为当前文件名

```

```bash 
    b main.go:10    设置断点，如果项目中存在同名文件，会根据行号对应的内容选择一个，或者两个都选
    run             运行到断点处
    n               运行到下一行
    ni              回到当前运行的代码片段
    p varname       打印变量值
    frame variable  打印当然 frame 中的所有变量
    r               重新运行程序
    x         -- ('memory read')  Read from the memory of the process being debugged

    LLDB treats Goroutines as threads.
    thread list
    bt all
    thread select 2
```

`root@ubuntu-xenial:/data/app/go/src/go1.11.1/demo/debug/lldb# go build -gcflags "-N -l" -o main main.go`
`root@ubuntu-xenial:/data/app/go/src/go1.11.1/demo/debug/lldb# lldb main`
`(lldb) target create "main"`
Current executable set to 'main' (x86_64).
`(lldb) b main.go:4`
Breakpoint 1: where = main`main.main + 29 at main.go:4, address = 0x000000000044c16d
`(lldb) r`
Process 3380 launched: '/data/app/go/src/go1.11.1/demo/debug/lldb/main' (x86_64)
Process 3380 stopped
* thread #1: tid = 3380, 0x000000000044c16d main`main.main + 29 at main.go:4, name = 'main', stop reason = breakpoint 1.1
    frame #0: 0x000000000044c16d main`main.main + 29 at main.go:4
   1   	package main
   2
   3   	func main() {
-> 4   		var i int
   5   		i = 100
   6   		println(i)
   7   	}
(lldb) n
Process 3380 stopped
* thread #1: tid = 3380, 0x000000000044c176 main`main.main + 38 at main.go:5, name = 'main', stop reason = step over
    frame #0: 0x000000000044c176 main`main.main + 38 at main.go:5
   2
   3   	func main() {
   4   		var i int
-> 5   		i = 100
   6   		println(i)
   7   	}
`(lldb) p i`
(int) i = 0
`(lldb) n`
Process 3380 stopped
* thread #1: tid = 3380, 0x000000000044c17f main`main.main + 47 at main.go:6, name = 'main', stop reason = step over
    frame #0: 0x000000000044c17f main`main.main + 47 at main.go:6
   3   	func main() {
   4   		var i int
   5   		i = 100
-> 6   		println(i)
   7   	}
`(lldb) p i`
(int) i = 100



`root@ubuntu-xenial:/data/app/go/src/go1.11.1/demo/debug/lldb# lldb lldb.test`
`(lldb) target create "lldb.test"`
Current executable set to 'lldb.test' (x86_64).
`(lldb) break set -r lldb.TestMain$`
Breakpoint 1: where = lldb.test`go1.11.1/demo/debug/lldb.TestMain + 29 at main_test.go:6, address = 0x00000000004e5f3d
`(lldb) run --test.run=TestMain`
Process 3701 launched: '/data/app/go/src/go1.11.1/demo/debug/lldb/lldb.test' (x86_64)
Process 3701 stopped
* thread #1: tid = 3701, 0x00000000004e5f3d lldb.test`go1.11.1/demo/debug/lldb.TestMain(t=0x000000c4200b40f0) + 29 at main_test.go:6, name = 'lldb.test', stop reason = breakpoint 1.1
    frame #0: 0x00000000004e5f3d lldb.test`go1.11.1/demo/debug/lldb.TestMain(t=0x000000c4200b40f0) + 29 at main_test.go:6
   3   	import "testing"
   4
   5   	func TestMain(t *testing.T) {
-> 6   		var i int
   7   		i = 100
   8   		println(i)
   9   	}
`(lldb) n`
Process 3701 stopped
* thread #1: tid = 3701, 0x00000000004e5f46 lldb.test`go1.11.1/demo/debug/lldb.TestMain(t=0x000000c4200b40f0) + 38 at main_test.go:7, name = 'lldb.test', stop reason = step over
    frame #0: 0x00000000004e5f46 lldb.test`go1.11.1/demo/debug/lldb.TestMain(t=0x000000c4200b40f0) + 38 at main_test.go:7
   4
   5   	func TestMain(t *testing.T) {
   6   		var i int
-> 7   		i = 100
   8   		println(i)
   9   	}
`(lldb) p i`
(int) i = 0
`(lldb) n`
Process 3701 stopped
* thread #1: tid = 3701, 0x00000000004e5f4f lldb.test`go1.11.1/demo/debug/lldb.TestMain(t=0x000000c4200b40f0) + 47 at main_test.go:8, name = 'lldb.test', stop reason = step over
    frame #0: 0x00000000004e5f4f lldb.test`go1.11.1/demo/debug/lldb.TestMain(t=0x000000c4200b40f0) + 47 at main_test.go:8
   5   	func TestMain(t *testing.T) {
   6   		var i int
   7   		i = 100
-> 8   		println(i)
   9   	}
(lldb) p i
(int) i = 100
`(lldb) si`
Process 3701 stopped
* thread #1: tid = 3701, 0x0000000000429340 lldb.test`runtime.printlock at print.go:66, name = 'lldb.test', stop reason = instruction step into
    frame #0: 0x0000000000429340 lldb.test`runtime.printlock at print.go:66
   63  	// the print lock to print information about the crash.
   64  	// For both these reasons, let a thread acquire the printlock 'recursively'.
   65
-> 66  	func printlock() {
   67  		mp := getg().m
   68  		mp.locks++ // do not reschedule between printlock++ and lock(&debuglock).
   69  		mp.printlock++
(lldb)