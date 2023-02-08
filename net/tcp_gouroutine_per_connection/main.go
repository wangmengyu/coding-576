package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func ListenAndServe(addr string) {

	listen, err := net.Listen("tcp", addr)
	defer listen.Close()
	if err != nil {
		panic(err.Error())
	}

	// 为每个请求设置自己的goroutie
	for {
		// 持续接收请求
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err.Error())
			continue
		}

		// 让每个链接在自己的协程里工作
		go handler(conn)
	}

}

func handler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		// 持续处理读写操作
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("conn closed")
			} else {
				fmt.Println("read err:", err.Error())
			}
			return
		}
		// 将数据写回通道
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("write err", err.Error())
			return
		}

	}

}

func main() {
	ListenAndServe(":8888")
}
