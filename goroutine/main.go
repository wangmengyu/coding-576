package main

import (
	"fmt"
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

	select {}

}
