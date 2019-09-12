package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

//服务端实现流程大致如下：
//
//创建监听listener，程序结束时关闭。
//阻塞等待客户端连接，程序结束时关闭conn。
//读取客户端发送文件名。保存fileName。
//回发“ok”给客户端做应答
//封装函数 RecvFile接收客户端发送的文件内容。传参fileName 和conn
//按文件名Create文件，结束时Close
//循环Read客户端发送的文件内容，当读到EOF说明文件读取完毕。
//将读到的内容原封不动Write到创建的文件中

func main() {
	//创建tcp监听
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	for {
		//阻塞等待客户端
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//创建协程
		go Handler(conn)
	}
}

func Handler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 2048)
	//读取客户端发送的内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileName := string(buf[:n])
	//获取客户端ip+port
	addr := conn.RemoteAddr().String()
	fmt.Println(addr + ": 客户端传输的文件名为--" + fileName)
	//告诉客户端已经接收到文件名
	conn.Write([]byte("ok"))

	//创建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//循环接收客户端传递的文件内容
	for {
		buf := make([]byte, 2048)
		n, _ := conn.Read(buf)
		//结束协程
		if string(buf[:n]) == "finish" {
			fmt.Println(addr + ": 协程结束")
			runtime.Goexit()
		}
		f.Write(buf[:n])
	}
}


// maxSendfileSize is the largest chunk size we ask the kernel to copy
// at a time.
const maxSendfileSize int = 4 << 20

// SendFile wraps the sendfile system call.
//func SendFile(dstFD *FD, src int, remain int64) (int64, error) {
//	if err := dstFD.writeLock(); err != nil {
//		return 0, err
//	}
//	defer dstFD.writeUnlock()
//
//	dst := int(dstFD.Sysfd)
//	var written int64
//	var err error
//	for remain > 0 {
//		n := maxSendfileSize
//		if int64(n) > remain {
//			n = int(remain)
//		}
//		n, err1 := syscall.Sendfile(dst, src, nil, n)
//		if n > 0 {
//			written += int64(n)
//			remain -= int64(n)
//		}
//		if n == 0 && err1 == nil {
//			break
//		}
//		if err1 == syscall.EAGAIN {
//			if err1 = dstFD.pd.waitWrite(dstFD.isFile); err1 == nil {
//				continue
//			}
//		}
//		if err1 != nil {
//			// This includes syscall.ENOSYS (no kernel
//			// support) and syscall.EINVAL (fd types which
//			// don't implement sendfile)
//			err = err1
//			break
//		}
//	}
//	return written, err
//}
