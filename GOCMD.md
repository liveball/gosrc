

// $GODEV/bin/go 使用开发模式
// Gosched 让出cpu
// Goexit 退出当前goroutine(但是defer语句会照常执行)
// GOMAXPROCS  设置逻辑处理器数量，即最大的可同时使用的CPU核数 set env: export GOMAXPROCS=1
// NumCPU 返回当前系统的CPU核数量

##  build
### go build -o main -gcflags "all=-N -l"

### 打印编译信息
  `go tool compile -d=slice,append,gcprog,closure,export,wb -o t.a main.go`

## print asm

### go tool objdump -S -s "main\.main" main

### go tool compile -N -l -S main.go > main.s

### go build -gcflags -S main.go  more asm

## gctrace
### go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main 

##  goroutine schedtrace
### go build -o main demo/runtime/goroutine/main.go && GODEBUG=schedtrace=10000,scheddetail=1 ./main


## test
### go tool compile -N -l -S demo/type/interface/struct/main.go > main.s

### go build -o main -gcflags "all=-N -l" demo/type/interface/struct/main.go && go tool objdump -S -s "main\.main" main > main.s

### list release tag
### go list  -f '{{context.ReleaseTags}}' runtime | grep go1.10
 [go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11]

go build -o main -gcflags "all=-N -l -m" demo/runtime/memory/closure/main.go
