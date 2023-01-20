package main

import (
	"fmt"
	"time"
)

func add(p *int32) {
	*p = *p + 1 // 非原子操作,会导致协程并发导致缺少 , 并发时不能这么写.
}

func main() {
	var i int32
	for j := 1; j <= 1000; j++ {
		go add(&i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(i)
}
