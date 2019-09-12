package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

//客户端实现流程大致如下：
//
//提示用户输入文件名。接收文件名path（含访问路径）
//使用os.Stat()获取文件属性，得到纯文件名（去除访问路径）
//主动连接服务器，结束时关闭连接
//给接收端（服务器）发送文件名conn.Write()
//读取接收端回发的确认数据conn.Read()
//判断是否为“ok”。如果是，封装函数SendFile() 发送文件内容。传参path和conn
//只读Open文件, 结束时Close文件
//循环读文件，读到EOF终止文件读取
//将读到的内容原封不动Write给接收端（服务器）

func main() {
	fmt.Print("请输入文件的完整路径：")

	//创建切片，用于存储输入的路径
	var path string

	fmt.Scan(&path)
	//获取文件信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	//创建客户端连接
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	//文件名称
	fileName := fileInfo.Name()
	//文件大小
	fileSize := fileInfo.Size()
	//发送文件名称到服务端
	conn.Write([]byte(fileName))

	buf := make([]byte, 2048)
	//读取服务端内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf[:n])
	if revData == "ok" {
		// /Users/fpf/Downloads/Linux1.0核心游记.pdf 6m
		// /Users/fpf/Downloads/深入理解Linux网络技术内幕.pdf 71m

		// 文件大小              函数              耗时

		//  6m                 ReadWriteFile    375.42198ms

		//  6m                 SendFile         322.495082ms

		//  71m                ReadWriteFile    4.592790698s

		//  71m                SendFile         2.396248848s

		//发送文件数据
		now := time.Now()
		//ReadWriteFile(path, fileSize, conn)

		SendFile(path, fileSize, conn, conn)
		fmt.Println("client 上传文件耗时:", time.Now().Sub(now))
	}
}

//发送文件到服务端
func ReadWriteFile(filePath string, fileSize int64, conn net.Conn) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	//var offset int64
	for {
		buf := make([]byte, 2048)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Println("文件传输完成")
			//告诉服务端结束文件接收
			conn.Write([]byte("finish"))
			return
		}
		//发送给服务端
		conn.Write(buf[:n])

		//offset += int64(n)
		//sendPercent := float64(offset) / float64(fileSize) * 100
		//value := fmt.Sprintf("%.2f", sendPercent)
		////打印上传进度
		//fmt.Println("文件上传：" + value + "%")
	}
}

//func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	return syscall.Sendfile(outfd, infd, offset, count)
}

//发送文件到服务端  /Users/fpf/Downloads/Linux1.0核心游记.pdf
func SendFile(filePath string, fileSize int64, c1, c2 net.Conn) {
	//f, err := os.Open(filePath)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//defer f.Close()
	//
	//infd:=int(f.Fd())//读入

	getFd := func(conn net.Conn) int {
		switch pconn := conn.(type) {
		case *net.TCPConn:
			cf, err := pconn.File()
			if err != nil {
				log.Printf("SendFile pconn.File error(%v)", err)
				return 0
			}
			return int(cf.Fd())
		}
		return 0
	}

	var outfd, infd int
	outfd = getFd(c1)
	infd = getFd(c2)

	if outfd == 0 || infd == 0 {
		return
	}

	var offset *int64
	offset = new(int64)
	for {

		written, err := sendfile(outfd, infd, offset, int(fileSize))
		if (written == int(fileSize)) || (err != nil && io.EOF == err) {

			//println(outfd, infd, *offset, int(fileSize))
			//sendPercent := float64(*offset) / float64(fileSize) * 100
			//value := fmt.Sprintf("%.2f", sendPercent)
			////打印上传进度
			//fmt.Println("文件上传：" + value + "%")

			fmt.Println("文件传输完成")
			c1.Write([]byte("finish"))
			return
		}
	}
}
