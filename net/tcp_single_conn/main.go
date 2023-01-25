package main

import (
	"fmt"
	"net"
)

func main() {
	// 本例用的单个协程来完成监听,建立连接.只为一个客户端服务.通常不这么使用.
	// 主要还是看tcp_goroutine_per_connection目录下的实现
	// 监听本地的8888端口 , 等待新的连接进来
	listen, err := net.Listen("tcp", ":8888")
	// listen的值是一个listen状态的socket.
	if err != nil {
		panic(err.Error())
	}
	conn, err := listen.Accept() // 接收新连接
	if err != nil {
		panic(err.Error())
	}

	// 新建缓存变量接收数据
	var body [100]byte
	// 以下这种方式一旦客户端断开连接. server就退出了.
	for {
		// 不断接受数据
		_, err = conn.Read(body[:])
		if err != nil {
			fmt.Println("read err:", err.Error())
			break // 接收失败退出本次循环
		}
		// 成功接收到数据 , 发回数据给客户端.
		fmt.Printf("收到数据: %s\n", body[:])
		_, err = conn.Write(body[:])
		if err != nil {
			fmt.Println("write err:", err.Error())
			break
		}
	}

}
