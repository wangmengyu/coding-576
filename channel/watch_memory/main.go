package main

import (
	"fmt"
	"time"
)

func watch(i *int) {
	for {
		if *i == 1 {
			fmt.Println("i==1")
			return
		}
	}
}

// 传统的监听变量做法。
func main() {
	i := 0
	go watch(&i)
	time.Sleep(1 * time.Second) //
	i = 1
	time.Sleep(1 * time.Second)
}
