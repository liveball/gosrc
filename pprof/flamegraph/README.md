
### go pprof 工具分析 CPU Heap

go tool pprof http://127.0.0.1:2333

go tool pprof -seconds=10 -pdf  http://127.0.0.1:2333/debug/pprof/profile > profile.pdf

go tool pprof -seconds=10 -pdf  http://127.0.0.1:2333/debug/pprof/heap > heap.pdf


### 生成火焰图
 
下载perl脚本
git clone https://github.com/brendangregg/FlameGraph

go tool pprof -seconds=10 -raw -output=a.pprof http://127.0.0.1:2333/debug/pprof/profile

./stackcollapse-go.pl a.pprof > pprof.folded  

./flamegraph.pl pprof.folded > pprof.svg


### 使用perf

export TERM=linux
apt-get update
apt-get install linux-tools-3.16

perf record

1. perf record -e cpu-clock -p PID

2. 把go binary 和 perf.data 放在一起，perf report -i perf.data