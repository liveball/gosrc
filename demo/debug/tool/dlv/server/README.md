
https://studygolang.com/articles/17961

https://blog.gopheracademy.com/advent-2018/postmortem-debugging-delve/

# 编译
go build -o server -gcflags "all=-N -l"

# 编译
go build -o server -gcflags "-N -l"

# 启动
nohup ./server --addr=:10080 > 1.log 2>&1 &

# 请求测试
curl -i 'http://127.0.0.1:10080'

curl -i --max-time 5 'http://127.0.0.1:10080'

# 压力测试
wrk -d1m -t4 -c1000 'http://127.0.0.1:10080'

hey -z 1m -n 1000 -c 1000 -t 60 -m GET 'http://127.0.0.1:10080'

# dlv core
gcore 5306 (PID)

dlv core ./server core.5306

https://rakyll.org/coredumps/

# dlv 远程调试
dlv core ./server core.5306 --listen :44441 --headless --log

echo -n '{"method":"RPCServer.ListGoroutines","params":[],"id":2}' | nc -w 1 localhost 44441 > list_goroutines.json


# jq 查询
### 查看3个对象
jq '.result[0:3]' list_goroutines.json

### 找到数量最多的goroutine

jq -c '.result[] | [.userCurrentLoc.function.name, .userCurrentLoc.line]' list_goroutines.json | sort | uniq -c

### 通过分析定位阻塞函数
jq '.result[] | select(.startLoc.function.name | test("startInChannelConsumer$"))' list_goroutines.json

### 切换goroutine 打印堆栈信息
dlv core ./server core.5306
goroutine 20
stack -full