package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func add(p *int32) {
	//*p = *p + 1 // 非原子操作,会导致协程并发导致缺少 , 并发时不能这么写.
	atomic.AddInt32(p, 1) // 用atomic的函数就可以保证原子性
}

func main() {
	var i int32
	for j := 1; j <= 1000; j++ {
		go add(&i)
	}
	atomic.CompareAndSwapInt32(&i, 10, 100)

	// 相当于
	if i == 10 {
		i = 100
	}

	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
