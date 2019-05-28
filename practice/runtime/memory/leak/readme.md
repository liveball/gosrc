### issue memory leak involving finalizers
https://github.com/golang/go/issues/631

### gc trace

`go build -gcflags "-N -l"  -o main /data/app/go/src/gosrc/practice/runtime/memory/leak/chan/main.go && GODEBUG=gctrace=1   ./main`

`go build -gcflags "-N -l"  -o main /data/app/go/src/gosrc/practice/runtime/memory/leak/fd/main.go && GODEBUG=gctrace=1   ./main`


### 可视化查看内存申请和释放
`get get -u -v github.com/davecheney/gcvis`
`GODEBUG=gctrace=1 ./main -index -http=:6060 2>&1 | gcvis`