package main

import (
	"fmt"
	"reflect"
)

func add1(a, b int) int {
	return a + b
}

func add2(a, b int) int {
	return -(a + b)
}

// CallAdd 外部只要调用这一个方法, 可以调用不同的实现, 传具体方法名就可以.
func CallAdd(f func(a, b int) int) {

	// 首先排除 f 不是func的情况
	fVal := reflect.ValueOf(f) // 获得f真正的值出来
	if fVal.Kind() != reflect.Func {
		fmt.Println("不是function")
		return
	}

	//设置好参数值: 两个value类型的参数
	args := make([]reflect.Value, 0)
	args = append(args, reflect.ValueOf(1))
	args = append(args, reflect.ValueOf(2))
	res := fVal.Call(args)
	fmt.Println(res[0])

}
func main() {
	CallAdd(add1)
	CallAdd(add2)

}
