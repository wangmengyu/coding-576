package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"github.com/wangmengyu/coding-576/utils"
	"unsafe"
)

// K 定义一个空的struct
type K struct {
}

func main() {
	fmt.Println("hello world")
	pool := tunny.Pool{}
	fmt.Printf("%+v, %T", pool, pool)
	val := utils.GetCurrentDate()
	fmt.Println("current date: ", val)

	//打印 Int 类型的字节数
	sizeOfInt := unsafe.Sizeof(int(0))
	fmt.Println("size of int:", sizeOfInt)

	i := int(0)
	p := &i
	fmt.Println("size of pointer of a int:", unsafe.Sizeof(p))

	f1 := float64(0)
	pf := &f1
	fmt.Println("size of point of a float64:", unsafe.Sizeof(pf))

	k1 := K{}
	n1 := int(0)
	k2 := K{}
	//pk1 := &k1>
	// 空对象的变量 没有占用任何产固定
	fmt.Println("size of point of a empty struct", unsafe.Sizeof(k1)) //
	// 但是它的地址: 有值
	fmt.Printf("address of empty struct K1: %p \n", &k1) // 十六进制,zero base, 位于rutime.malloc文件中可以查到
	fmt.Printf("address of empty struct: %p \n", &n1)    // 十六进制,有具体地址
	fmt.Printf("address of empty struct K2: %p \n", &k2) // 十六进制,zero base

	m := map[string]struct{}{}
	m["1"] = struct{}{}
	fmt.Printf("size of m-1: %d\n", unsafe.Sizeof(m["1"]))

	ch1 := make(chan struct{})
	fmt.Printf("size of ch1-1: %d\n", unsafe.Sizeof(ch1))

}
