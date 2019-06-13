### go 链接器

### 测试程序(https://github.com/liveball/gosrc/tree/master/practice/link/stack)
``
>1.开启godev命令模式(https://github.com/liveball/gosrc/blob/master/README.md)
>2.编译 .o 文件
`go tool compile -o main.o /data/app/go/src/gosrc/practice/link/stack/main.go`
>3.链接 .o 文件
`go run /data/app/go/src/gosrc/go/src/cmd/link/main.go -o main main.o`


###分裂栈检查
go/src/cmd/link/internal/ld/lib.go

```go 
   func (ctxt *Link) dostkcheck() 
```

### 栈的大小限制
go/src/cmd/internal/objabi/stack.go
```go
package objabi

// For the linkers. Must match Go definitions.

const (
	STACKSYSTEM = 0
	StackSystem = STACKSYSTEM
	StackBig    = 4096
	StackSmall  = 128
)

const (
	StackPreempt = -1314 // 0xfff...fade
)

// Initialize StackGuard and StackLimit according to target system.
var StackGuard = 880*stackGuardMultiplier() + StackSystem
var StackLimit = StackGuard - StackSystem - StackSmall

// stackGuardMultiplier returns a multiplier to apply to the default
// stack guard size. Larger multipliers are used for non-optimized
// builds that have larger stack frames or for specific targets.
func stackGuardMultiplier() int {
	// On AIX, a larger stack is needed for syscalls.
	if GOOS == "aix" {
		return 2
	}
	return stackGuardMultiplierDefault
}
```
