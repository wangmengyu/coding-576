package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"github.com/wangmengyu/coding-576/utils"
)

func main() {
	fmt.Println("hello world")
	pool := tunny.Pool{}
	fmt.Printf("%+v, %T", pool, pool)
	val := utils.GetCurrentDate()
	fmt.Println("current date: ", val)
}
