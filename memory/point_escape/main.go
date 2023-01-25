package main

import "fmt"

// 指针逃逸范例:
func a() *int {
	v := 0
	return &v // 返回一个局部变量的指针. 发生指针逃逸.这种情况下v的指针被放在堆上. 不会被回收.

}
func main() {
	i := a() // 接收a方法的返回的指针
	fmt.Println(i)
}
