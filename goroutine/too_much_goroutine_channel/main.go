package main

import (
	"fmt"
	"math"
	"time"
)

func do1(i int, ch chan struct{}) {
	defer func() {
		<-ch
	}()
	fmt.Println(i)
	time.Sleep(1 * time.Second)
}
func main() {
	// 利用channel 让 大量 goroutine 可以执行. 保证同时执行的goroutine  在一定的数量
	ch := make(chan struct{}, 3000)
	for i := 0; i < math.MaxInt32; i++ {
		// 开启每个协程之前, 先给通道释放一个信号
		ch <- struct{}{}
		go do1(i, ch)
	}
	select {}

}
