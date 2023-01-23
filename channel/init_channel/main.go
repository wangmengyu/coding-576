package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1) // 声明channel ,没有缓冲
	ch <- "1"                  // 送入一个值
	go GetData(ch)
	time.Sleep(1 * time.Second)
}

func GetData(ch chan string) {
	v := <-ch
	fmt.Println(v)
}
