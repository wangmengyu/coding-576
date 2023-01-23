package main

import "fmt"

// 非阻塞式channel
func main() {
	c1 := make(chan int, 5) // 有5个缓存
	c2 := make(chan int)    // 没有缓存

	// 如果在这里直接写 <-c1 是会卡主的.因为没有数据可以被取出
	// 如果在这列直接写 c2<-1  也会卡主, 因为没有人进行消耗
	select { // 但是写select 里面就不会被卡住.
	case <-c1: // 如果从c1接收到了数据.跑这里
		fmt.Println("c1")
	case c2 <- 1: //  如果想C2成功发送了数据. 跑这里
		fmt.Println("c2")
	default: // 其他情况
		fmt.Println("none")
	}

}
