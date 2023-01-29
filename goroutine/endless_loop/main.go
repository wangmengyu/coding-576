package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// 无函数调用.
				a[i]++
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("I got scheduled:", a)
	// 运行结果能返回到I got scheduled , Go 1.14 引入了基于系统信号的异步抢占调度，这样，像上面的无函数调用的死循环
	//goroutine 也可以被抢占了，从而将 main goroutine 重新调度回 P 执行，最终会打印 I got scheduled!。

}
