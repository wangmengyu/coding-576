package main

import "fmt"

func main() {
	// 字面量创建, 编译时执行
	s1 := []int{1, 2, 3}
	fmt.Println(s1)

	// 切片的创建
	a1 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := a1[1:4]   // 截取方式是:的去左边计算, 右边不计算
	fmt.Println(s2) // [1,2,3]

	// 切片的扩容
	s2 = append(s2, 8, 10, 12)
	fmt.Println(s2)

	// 删除第4个元素
	s3 := append(s2[:3], s2[4:]...)
	fmt.Println(s3)

}
