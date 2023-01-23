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

func watch(ch chan int) { //  不需要循环.因为管道一直被监听.
	if <-ch == 1 {
		fmt.Println("it is one")
	}
}
