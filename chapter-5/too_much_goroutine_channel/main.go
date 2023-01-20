package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int, ch chan struct{}) {
	fmt.Println(i)
	time.Sleep(1 * time.Second)
	//在处理玩业务逻辑后是放掉一个channel的一个元素
	<-ch

}

// 使用channel 控制
func main() {
	// 定义一个3000缓存的channel
	ch := make(chan struct{}, 3000)
	for i := 0; i < math.MaxInt32; i++ {
		// 在开启goroutine之前, 先释放一个信号到通道
		ch <- struct{}{}
		go do(i, ch)
	}
	select {}
}
