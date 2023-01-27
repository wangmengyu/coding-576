package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 查看变量所占字节大小
	fmt.Println(unsafe.Sizeof(0))

}
