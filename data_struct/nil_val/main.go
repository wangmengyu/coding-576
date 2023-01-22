package main

import "fmt"

func main() {
	// nil 为空指针时
	var a *int
	fmt.Println(a == nil) // true
	var b map[int]int
	fmt.Println(b == nil) //  true

	var c interface{} // 空接口类型
	var d interface{} // 空接口类型

	fmt.Println(c == nil) // true
	fmt.Println(d == nil) //true

	// 如果将a赋值给c 此时c不再是空接口.因为c的Type被赋值成a的类型了. 就不再是nil
	c = a
	fmt.Println(c == nil)       // false
	fmt.Printf("%+v, %T", c, c) //<nil>, *int

}
