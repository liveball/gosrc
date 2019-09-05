### go 内存管理分析

### Proposal: Eliminate STW stack re-scanning
https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md

### Allocation efficiency in high-performance Go services
https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/


### maps 查看进程的内存区域映射信息
其中注意的一点是[stack:<tid>]是线程的堆栈信息，对应于/proc/[pid]/task/[tid]/路径。

```cmd
sudo cat /proc/15287/maps |grep "\<c0[0-9a-f]"

c000000000-c000200000 rw-p 00000000 00:00 0 
c000200000-c004000000 rw-p 00000000 00:00 0 
```

### strace 进程追踪内存分配情况

```cmd
strace -e mmap,mmap2,munmap,brk,mremap,remap_file_pages -kfp 15287

strace: Process 15287 attached with 7 threads
[pid 15289] mmap(NULL, 8392704, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_STACK, -1, 0) = 0x7f17de7fd000
 > /lib/x86_64-linux-gnu/libc-2.27.so(mmap64+0x43) [0x11ba13]
 > /lib/x86_64-linux-gnu/libpthread-2.27.so(pthread_create+0x766) [0x8116]
 > /tmp/go-build761524685/b001/exe/main() [0x12f9]
 > unexpected_backtracing_error [0x800000]
strace: Process 15302 attached
[pid 15302] mmap(NULL, 134217728, PROT_NONE, MAP_PRIVATE|MAP_ANONYMOUS|MAP_NORESERVE, -1, 0) = 0x7f17c0000000
 > /lib/x86_64-linux-gnu/libc-2.27.so(mmap64+0x43) [0x11ba13]
 > /lib/x86_64-linux-gnu/libc-2.27.so(_IO_str_seekoff+0x859) [0x90c89]
 > /lib/x86_64-linux-gnu/libc-2.27.so(_IO_str_seekoff+0x162c) [0x91a5c]
 > /lib/x86_64-linux-gnu/libc-2.27.so(_IO_str_seekoff+0x611d) [0x9654d]
 > /lib/x86_64-linux-gnu/libc-2.27.so(cfree+0x54e) [0x97e9e]
 > /tmp/go-build761524685/b001/exe/main() [0x13da]
[pid 15302] munmap(0x7f17c4000000, 67108864) = 0
 > /lib/x86_64-linux-gnu/libc-2.27.so(__munmap+0x7) [0x11bab7]
 > /lib/x86_64-linux-gnu/libc-2.27.so(_IO_str_seekoff+0x897) [0x90cc7]
 > /lib/x86_64-linux-gnu/libc-2.27.so(_IO_str_seekoff+0x162c) [0x91a5c]
 > /lib/x86_64-linux-gnu/libc-2.27.so(_IO_str_seekoff+0x611d) [0x9654d]
 > /lib/x86_64-linux-gnu/libc-2.27.so(cfree+0x54e) [0x97e9e]
 > /tmp/go-build761524685/b001/exe/main() [0x13da]
 
 ```

### limits /proc/[pid]/limits显示当前进程的资源限制
    `cat /proc/$pid/limits`
    
### stack /proc/[pid]/stack显示当前进程的内核调用栈信息，

    只有内核编译时打开了CONFIG_STACKTRACE编译选项，才会生成这个文件。举例如下：

```
sudo cat /proc/15287/task/15287/stack

[<0>] futex_wait_queue_me+0xc4/0x120
[<0>] futex_wait+0x10a/0x250
[<0>] do_futex+0x364/0x570
[<0>] __x64_sys_futex+0x13f/0x190
[<0>] do_syscall_64+0x5a/0x120
[<0>] entry_SYSCALL_64_after_hwframe+0x44/0xa9
[<0>] 0xffffffffffffffff
```

### statm /proc/[pid]/statm显示进程所占用内存大小的统计信息，
包含七个值，度量单位是page（page大小可通过getconf PAGESIZE得到）
```
 cat /proc/2948/statm
 72362 12945 4876 569 0 24665 0
```
各个值含义：
a）进程占用的总的内存；
b）进程当前时刻占用的物理内存；
c）同其它进程共享的内存；
d）进程的代码段；
e）共享库（从2.6版本起，这个值为0）；
f）进程的堆栈；
g）dirty pages（从2.6版本起，这个值为0）。

### syscall /proc/[pid]/syscall显示当前进程正在执行的系统调用

```cmd
cat /proc/2948/syscall
7 0x7f4a452cbe70 0xb 0x1388 0xffffffffffdff000 0x7f4a4274a750 0x0 0x7ffd1a8033f0 0x7f4a41ff2c1d
```

第一个值是系统调用号（7代表poll），后面跟着6个系统调用的参数值（位于寄存器中），最后两个值依次是堆栈指针和指令计数器的值。
如果当前进程虽然阻塞，但阻塞函数并不是系统调用，则系统调用号的值为-1，后面只有堆栈指针和指令计数器的值。
如果进程没有阻塞，则这个文件只有一个“running”的字符串。

内核编译时打开了CONFIG_HAVE_ARCH_TRACEHOOK编译选项，才会生成这个文件。

### wchan /proc/[pid]/wchan显示当进程sleep时，kernel当前运行的函数

```cmd
 cat /proc/2948/wchan
 futex_wait_queue_me_
```

### root /proc/[pid]/root是进程根目录的符号链接
```cmd
 ls -lt /proc/15287/root
 lrwxrwxrwx 1 root root 0 9月   4 15:37 /proc/15287/root -> /
```

### auxv /proc/[pid]/auxv包含传递给进程的ELF解释器信息

格式是每一项都是一个unsigned long长度的ID加上一个unsigned long长度的值。最后一项以连续的两个0x00开头。举例如下：

```cmd # hexdump -x /proc/2948/auxv
0000000    0021    0000    0000    0000    0000    1a82    7ffd    0000
0000010    0010    0000    0000    0000    dbf5    1fc9    0000    0000
0000020    0006    0000    0000    0000    1000    0000    0000    0000
0000030    0011    0000    0000    0000    0064    0000    0000    0000
0000040    0003    0000    0000    0000    2040    4326    7f4a    0000
0000050    0004    0000    0000    0000    0038    0000    0000    0000
0000060    0005    0000    0000    0000    0009    0000    0000    0000
0000070    0007    0000    0000    0000    f000    4303    7f4a    0000
0000080    0008    0000    0000    0000    0000    0000    0000    0000
0000090    0009    0000    0000    0000    8e67    4327    7f4a    0000
00000a0    000b    0000    0000    0000    0000    0000    0000    0000
00000b0    000c    0000    0000    0000    0000    0000    0000    0000
00000c0    000d    0000    0000    0000    0000    0000    0000    0000
00000d0    000e    0000    0000    0000    0000    0000    0000    0000
00000e0    0017    0000    0000    0000    0000    0000    0000    0000
00000f0    0019    0000    0000    0000    3de9    1a80    7ffd    0000
0000100    001f    0000    0000    0000    4fe5    1a80    7ffd    0000
0000110    000f    0000    0000    0000    3df9    1a80    7ffd    0000
0000120    0000    0000    0000    0000    0000    0000    0000    0000
0000130
```
解析这个文件可以参考这段代码。

### cmdline /proc/[pid]/cmdline是一个只读文件，包含进程的完整命令行信息
如果这个进程是zombie进程，则这个文件没有任何内容

```cmd # ps -ef | grep 2948
root       2948      1  0 Nov05 ?        00:00:04 /usr/sbin/libvirtd --listen

# cat /proc/2948/cmdline
/usr/sbin/libvirtd--listen
```
###comm /proc/[pid]/comm包含进程的命令名

```cmd # cat /proc/2948/comm
libvirtd
```

### cwd /proc/[pid]/cwd是进程当前工作目录的符号链接

```cmd # ls -lt /proc/2948/cwd
lrwxrwxrwx 1 root root 0 Nov  9 12:14 /proc/2948/cwd -> /
```

###environ /proc/[pid]/environ显示进程的环境变量

```cmd # strings /proc/2948/environ
LANG=POSIX
LC_CTYPE=en_US.UTF-8
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
NOTIFY_SOCKET=@/org/freedesktop/systemd1/notify
LIBVIRTD_CONFIG=/etc/libvirt/libvirtd.conf
LIBVIRTD_ARGS=--listen
LIBVIRTD_NOFILES_LIMIT=2048
```

### exe /proc/[pid]/exe为实际运行程序的符号链接

```cmd # ls -lt /proc/2948/exe
lrwxrwxrwx 1 root root 0 Nov  5 13:04 /proc/2948/exe -> /usr/sbin/libvirtd
```

### fd /proc/[pid]/fd是一个目录，包含进程打开文件的情况
```cmd
# root@fpf-VirtualBox:/home/fpf# ls -lt /proc/15287/fd
  总用量 0
  lrwx------ 1 root root 64 9月   4 16:05 0 -> /dev/pts/0
  lrwx------ 1 root root 64 9月   4 16:05 1 -> /dev/pts/0
  lrwx------ 1 root root 64 9月   4 16:05 2 -> /dev/pts/0
  lrwx------ 1 root root 64 9月   4 16:05 3 -> 'socket:[326501]'
  lrwx------ 1 root root 64 9月   4 16:05 5 -> 'anon_inode:[eventpoll]'
```
目录中的每一项都是一个符号链接，指向打开的文件，数字则代表文件描述符。

### latency /proc/[pid]/latency显示哪些代码造成的延时比较大（使用这个feature，需要执行“echo 1 > /proc/sys/kernel/latencytop”）
```cmd
# cat /proc/2948/latency
Latency Top version : v0.1
30667 10650491 4891 poll_schedule_timeout do_sys_poll SyS_poll system_call_fastpath 0x7f636573dc1d
8 105 44 futex_wait_queue_me futex_wait do_futex SyS_futex system_call_fastpath 0x7f6365a167bc
每一行前三个数字分别是后面代码执行的次数，总共执行延迟时间（单位是微秒）和最长执行延迟时间（单位是微秒），后面则是代码完整的调用栈。
```

### limits /proc/[pid]/limits显示当前进程的资源限制

``` cmd
#cat /proc/2948/limits
Limit                     Soft Limit           Hard Limit           Units
Max cpu time              unlimited            unlimited            seconds
Max file size             unlimited            unlimited            bytes
Max data size             unlimited            unlimited            bytes
Max stack size            8388608              unlimited            bytes
Max core file size        0                    unlimited            bytes
Max resident set          unlimited            unlimited            bytes
Max processes             6409                 6409                 processes
Max open files            1024                 4096                 files
Max locked memory         65536                65536                bytes
Max address space         unlimited            unlimited            bytes
Max file locks            unlimited            unlimited            locks
Max pending signals       6409                 6409                 signals
Max msgqueue size         819200               819200               bytes
Max nice priority         0                    0
Max realtime priority     0                    0
Max realtime timeout      unlimited            unlimited            us
Soft Limit表示kernel设置给资源的值，Hard Limit表示Soft Limit的上限，而Units则为计量单元。
```