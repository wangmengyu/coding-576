package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "一二三"
	// 将s转成stringHeader类型
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sh.Len) // 打印9
	for _, ch := range s {
		fmt.Printf("%c\n", ch)
	}

	// 字符串切分
	s2 := string([]rune(s)[:2])
	fmt.Println(s2)
}
