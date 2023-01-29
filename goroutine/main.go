package main

import (
	"fmt"
	"time"
)

// 定义3个方嵌套调用

func do1() {
	do2()

}
func do2() {

	do3()
}
func do3() {
	fmt.Println("do3")
}

func main() {
	// 开启协程, 执行do1
	go do1()

	time.Sleep(1 * time.Second)

	fmt.Println("done")
}
