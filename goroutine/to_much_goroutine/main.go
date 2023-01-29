package main

import (
	"fmt"
	"math"
	"time"
)

func do(i int) {
	fmt.Println(i)
	time.Sleep(1 * time.Second)
}
func main() {
	// 起math.MaxInt32个协程执行方法
	for i := 0; i < math.MaxInt32; i++ {
		go do(i)
	}
	select {}
}
