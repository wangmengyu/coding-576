package main

import "fmt"

func b() {
	i := 0
	fmt.Println(i)
	// 这个方法接收的是...Interface{}, 很有可能会逃逸 ,
	// 原因是interface{}往往会使用到反射, 会看一下到底实参是什么类型.
	// 因为反射要求数据是在堆上的.

}
func main() {

}
