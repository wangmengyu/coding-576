package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

// ListenAndServe 每一个链接一个goroutine.
func ListenAndServe(address string) {
	listen, err := net.Listen("tcp", address)
	defer listen.Close() // 每次listen好, 就接defer关闭.
	if err != nil {
		// 碰到错误直接退出
		log.Fatal(err.Error())
	}

	// 不断接受新的请求, 创建出新的请求连接出来
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("accept err:", err.Error())
		}
		go Handle(conn)
	}
}

// Handle 处理每个通道的读写
func Handle(conn net.Conn) {
	reader := bufio.NewReader(conn) // 使用bufio标准库提供的缓冲区功能充当读取器
	for {
		// ReadString 会一直阻塞直到遇到分隔符 '\n'
		// 遇到分隔符后会返回上次遇到分隔符或连接建立后收到的所有数据, 包括分隔符本身
		// 若在遇到分隔符之前遇到异常, ReadString 会返回已收到的数据和错误信息
		msg, err := reader.ReadString('\n') //
		if err != nil {
			// 错误分为中断连接时候类型是io.EOF
			if err == io.EOF {
				log.Println("connection close")
			} else {
				log.Println("read msg err: ", err.Error())
			}
			return // 发生错误必须返回了
		}
		b := []byte(msg) // 强制转成[]byte类型
		fmt.Println("read msg:", string(b))
		_, err = conn.Write(b) // 将接收到的数据写回到连接中
		if err != nil {
			log.Println("write err:", err.Error())
			return
		}
	}
}

func main() {
	ListenAndServe(":8888")
}
