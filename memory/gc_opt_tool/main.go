package main

import "sync"

// 协程. 定义一个数字.把数字加到e的10次方

func addNum(wg *sync.WaitGroup) {
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go addNum(&wg)
	}
	wg.Wait()
}


