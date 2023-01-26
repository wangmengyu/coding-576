package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("defer main g")
	go func() {
		defer fmt.Println("defer g")
		panic("123")
		fmt.Println("end g")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end main g ")

}
