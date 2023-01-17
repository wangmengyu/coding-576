package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
)

func main() {
	fmt.Println("hello world")
	pool := tunny.Pool{}
	fmt.Printf("%+v, %T", pool, pool)
}
