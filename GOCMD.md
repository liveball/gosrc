

// $GODEV/bin/go 使用开发模式
// Gosched 让出cpu
// Goexit 退出当前goroutine(但是defer语句会照常执行)
// GOMAXPROCS  设置逻辑处理器数量，即最大的可同时使用的CPU核数 set env: export GOMAXPROCS=1
// NumCPU 返回当前系统的CPU核数量

##  build
### go build -o main -gcflags "all=-N -l"

## print asm

### go tool objdump -S -s "main\.main" main

### go tool compile -N -l -S main.go > main.s

## gctrace
### go build -o main -gcflags "-N -l" && GODEBUG=gctrace=1   ./main 

##  goroutine schedtrace
### go build -o main demo/runtime/goroutine/main.go && GODEBUG=schedtrace=10000,scheddetail=1 ./main


## test
### go tool compile -N -l -S demo/type/interface/struct/main.go > main.s

### go build -o main -gcflags "all=-N -l" demo/type/interface/struct/main.go && go tool objdump -S -s "main\.main" main > main.s



go build -o main -gcflags "all=-N -l -m" demo/runtime/memory/closure/main.go
