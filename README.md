## 深入go1.11.1 源码学习

### 使用GODEV 跑标准库的test

>1、设置GODEV 为自己的源码目录  

` $  GODEV= go1.11.1 的绝对路径 `  

>2、进入源码目录

` $  cd $GODEV/src `  

>3、设置引导安装到当前安装目录

` $ GOROOT_BOOTSTRAP=/ ./all.bash ` use /  

` $ GOROOT_BOOTSTRAP=$(go env GOROOT) ./make.bash ` use GOROOT  


```

Building Go cmd/dist using /usr/local/go.
Building Go toolchain1 using /usr/local/go.
Building Go bootstrap cmd/go (go_bootstrap) using Go toolchain1.
Building Go toolchain2 using go_bootstrap and Go toolchain1.
Building Go toolchain3 using go_bootstrap and Go toolchain2.
Building packages and commands for darwin/amd64.
---
##### API check
Go version is "go1.11.1", ignoring -next /data/app/go/src/go1.11.1/api/next.txt

ALL TESTS PASSED
---
Installed Go for darwin/amd64 in /data/app/go/src/go1.11.1
Installed commands in /data/app/go/src/go1.11.1/bin
*** You need to add /data/app/go/src/go1.11.1/bin to your PATH.

```

>4、使用新编译的工具链运行所有测试  

 >4.1 使用mod作为包管理工具  
   `export GO111MODULE=on`

 >4.2 默认使用当前目录名作为当前包名的mod，如果需要重命名则删除存在的mod，重新生成  
   `$GODEV/bin/go mod init runtime`

>4.3、进入某个标准包下面运行单个测试  

   `$GODEV/bin/go test -v -run=TestYieldLocked`

    === RUN   TestYieldLockedProgress
	--- PASS: TestYieldLockedProgress (0.01s)
	=== RUN   TestYieldLocked
	--- PASS: TestYieldLocked (0.01s)
	PASS
	ok  	runtime	0.031s

   `$GODEV/bin/go test -run=NONE -v -bench=ChanProdCons0`
	goos: darwin
	goarch: amd64
	pkg: runtime
	BenchmarkChanProdCons0-4         3000000               471 ns/op
	PASS
	ok      runtime 1.919s

    `$GODEV/bin/go test -v -bench=. -run=Benchmark`
	goos: darwin
	goarch: amd64
	pkg: runtime
	BenchmarkMakeChan/Byte-4                        30000000                47.3 ns/op
	BenchmarkMakeChan/Int-4                         20000000                54.2 ns/op
	BenchmarkMakeChan/Ptr-4                         10000000               130 ns/op
	BenchmarkMakeChan/Struct/0-4                    30000000                41.3 ns/op
   

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("read go")
}

```