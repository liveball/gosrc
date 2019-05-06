package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	maxListenerBacklog()
}

func maxListenerBacklog() int {
	fd, err := os.Open("/proc/sys/net/core/somaxconn")
	if err != nil {
		return syscall.SOMAXCONN
	}
	defer fd.Close()

	rd := bufio.NewReader(fd)
	line, err := rd.ReadString('\n')
	fmt.Println("line", line)
	if err != nil {
		return syscall.SOMAXCONN
	}

	f := strings.Fields(line)
	if len(f) < 1 {
		return syscall.SOMAXCONN
	}

	n, err := strconv.Atoi(f[0])
	if err != nil || n == 0 {
		return syscall.SOMAXCONN
	}
	fmt.Println("f", f)
	println("n", n, "计算新的n值", 1<<16-1)
	// Linux stores the backlog in a uint16.
	// Truncate number to avoid wrapping.
	// See issue 5030.
	if n > 1<<16-1 {
		n = 1<<16 - 1
	}

	return n
}
