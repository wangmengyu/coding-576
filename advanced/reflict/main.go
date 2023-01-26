package main

import (
	"fmt"
	"reflect"
)

func main() {

	s := "kinopio"
	// 拿到s的类型
	sType := reflect.TypeOf(s)
	// 打印s类型的名字, align
	fmt.Println(sType.Name())
	fmt.Println(sType.Align())

	// 拿到s的值
	sVal := reflect.ValueOf(s)
	// 获得值的长度
	fmt.Println(sVal.Len())

	// 还原值到原来的变量
	s2 := sVal.Interface().(string)
	fmt.Println(s2)
}
