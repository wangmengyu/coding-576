package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		panic("123")
		fmt.Println("end g")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("end main g ")

}
