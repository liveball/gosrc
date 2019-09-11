package main

import (
	"fmt"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
)

var (
	pid = unix.Getpid()
)

func main() {
	fmt.Println("pid:", pid)

	if err := syscall.Kill(pid, 0); err == nil {
		fmt.Printf("立即退出进程 pid(%d) err(%v):", pid, err)
		return
	}

	time.Sleep(20 * time.Second)
}

func getid(s string) {
	pid := unix.Getpid()
	fmt.Println(s+" Getpid", pid)
	fmt.Println(s+" Getppid", unix.Getppid())

	pgid, _ := unix.Getpgid(pid)
	fmt.Println(s+" Getpgid", pgid)

	sid, _ := unix.Getsid(pid)
	fmt.Println(s+" Getsid", sid)

	fmt.Println(s+" Getegid", unix.Getegid())
	fmt.Println(s+" Geteuid", unix.Geteuid())
	fmt.Println(s+" Getgid", unix.Getgid())
	fmt.Println(s+" Getuid", unix.Getuid())
}

func KillRelease(pid int) bool {
	//if(pid >0){
	//	h, _ := syscall.OpenProcess(syscall.TOKEN_ALL_ACCESS, false, uint32(pid))
	//	err1 := syscall.TerminateProcess(h,0)
	//	if (err1 != nil){
	//		syscall.CloseHandle(h);
	//		return false
	//	}else {
	//		syscall.WaitForSingleObject(h, 2000); // At most ,waite 2000  millisecond.
	//	}
	//	syscall.CloseHandle(h);
	//	return true;
	//}
	return false
}
