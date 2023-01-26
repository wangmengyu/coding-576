package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("defer main g")
	}()
	go func() {

		defer func() {
			recover()
			fmt.Println("defer g")
		}()

		panic("123")
		fmt.Println("after 123")
		fmt.Println("end g")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end main g")
}
