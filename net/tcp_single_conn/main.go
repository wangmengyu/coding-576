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
	if err != nil {
		panic(err.Error())
	}

	conn, err := listen.Accept()
	if err != nil {

	}

	var body [100]byte
	for {
		// 不断尝试获得一个数据
		_, err = conn.Read(body[:]) // 读取到body 里
		if err != nil {
			fmt.Println("read err: ", err.Error())
			// 读取发生错误立即退出循环
			break
		}

		_, err = conn.Write(body[:]) // 把数据写回通道
		if err != nil {
			// 写发生错误 . 也立即退出

			fmt.Println("write err:", err.Error())
			break
		}

	}

}
