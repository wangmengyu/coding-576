package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 新建m
	m := sync.Map{}

	// 存 协程
	go func() {
		for i := 0; i < 10000; i++ {
			m.Store(i, i)
		}
	}()
	// 读 协程
	go func() {
		for i := 9999; i >= 0; i-- {
			value, ok := m.Load(i)
			if !ok {
				fmt.Printf("not found key: %d\n ", i)
				continue
			}
			fmt.Printf("get key val: [%d]=>[%d]\n", i, value)
		}
	}()

	time.Sleep(1 * time.Second)
}
