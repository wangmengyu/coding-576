package main

import (
	"fmt"
	"time"
)

func main() {
	// 需求, 监听管道内的数据.一旦出现1就退出运行
	ch := make(chan int)
	go watch(ch)
	time.Sleep(1 * time.Second)
	ch <- 1
	time.Sleep(1 * time.Second)

}

func watch(ch chan int) {
	if <-ch == 1 {
		fmt.Println("听到了1")
	}
}
