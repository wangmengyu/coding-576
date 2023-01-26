package main

import "fmt"

type company struct {
	name    string
	address string
	boss    *people
}
type people struct {
	name string
}

func main() {

	b1 := people{name: "kinopio"}
	c1 := &company{
		boss: &b1, // 二级引用, 不能被GC
	}

	fmt.Println(c1)
}
