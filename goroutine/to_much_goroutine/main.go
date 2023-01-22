package main

import (
	"fmt"
	"math"
	"time"
)

// 打印i并且停止一秒
func do(i int) {
	fmt.Println(i)
	time.Sleep(time.Second)
}
func main() {
	//起很多协程都执行do方法
	for i := 0; i < math.MaxInt32; i++ {
		go do(i)
	}
	select {}
}
