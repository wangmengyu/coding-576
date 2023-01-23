package main

import (
	"fmt"
	"sync"
)

// 使用 channel 开发一个基本的互斥锁
// 思路: channel 本身就是一种锁概念, 当被存入数据的时候 没有人消耗会阻塞.
// 当消耗的掉的时候.就解除了锁定.

// MyMutex 定义锁结构体, 一个空结构体的channel
type MyMutex chan struct{}

func (m *MyMutex) Lock() {
	(*m) <- struct{}{}
}

func (m *MyMutex) UnLock() {
	<-(*m)
}

// NewMyMutex 新建锁
func NewMyMutex() MyMutex {
	ch := make(chan struct{}, 1)
	return ch
}

type Data struct {
	mutex MyMutex
	val   int
}

func main() {
	// 使用锁
	d := &Data{
		mutex: NewMyMutex(),
		val:   0,
	}
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go d.addOne(&wg)
	}
	wg.Wait()
	fmt.Println("done")
}

func (d *Data) addOne(wg *sync.WaitGroup) {
	d.mutex.Lock()
	defer d.mutex.UnLock()
	d.val = d.val + 1
	fmt.Println(d.val)
	wg.Done()

}
