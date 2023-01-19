package main

import (
	"fmt"
	"unsafe"
)

type S1 struct {
	num1 int32 //4字节
	num2 int32 // 4字节
}
type S2 struct {
	num1 int16 // 2字节
	num2 int32 // 4字节
}

func main() {
	// s1的值占用的空间
	fmt.Println(unsafe.Sizeof(S1{})) // 8字节
	// s2占用的空间.
	fmt.Println(unsafe.Sizeof(S2{})) // 8字节.

	fmt.Printf("bool size of:%d align:%d \n", unsafe.Sizeof(true), unsafe.Alignof(true))
	fmt.Printf("byte size of:%d align:%d \n", unsafe.Sizeof(byte(0)), unsafe.Alignof(byte(0)))
	fmt.Printf("int8 size of:%d align:%d \n", unsafe.Sizeof(int8(0)), unsafe.Alignof(int8(0)))
	fmt.Printf("int16 size of:%d align:%d \n", unsafe.Sizeof(int16(0)), unsafe.Alignof(int16(0)))
	fmt.Printf("int32 size of:%d align:%d \n", unsafe.Sizeof(int32(0)), unsafe.Alignof(int32(0)))
	fmt.Printf("int64 size of:%d align:%d \n", unsafe.Sizeof(int64(0)), unsafe.Alignof(int64(0)))
}
