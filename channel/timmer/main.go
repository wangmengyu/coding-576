package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(time.Second) // 定时器
	/**
	type Timer struct {
		C <-chan Time // C会在 倒计时结束的时候塞入一个数据.
		r runtimeTimer
	}
	*/
	<-t.C
	fmt.Println("hello")
}
