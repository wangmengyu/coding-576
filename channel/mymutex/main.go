package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MyMutex 用channel实现一个Mutex
type MyMutex chan struct{}

func NewMutex() MyMutex {
	m := make(chan struct{}, 1) // 同时进行的协程只能有一个
	return m
}

func (m *MyMutex) Lock() {
	*m <- struct{}{}
}

func (m *MyMutex) UnLock() {
	<-*m
}

type DemoData struct {
	num  int
	lock MyMutex
}

func (dd *DemoData) Add(done chan struct{}) {
	dd.lock.Lock()
	defer dd.lock.UnLock() // 解锁
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(3000)))
	dd.num++
	fmt.Println(dd.num)
	done <- struct{}{}
}

func main() {

	dd := DemoData{
		num:  0,
		lock: NewMutex(),
	}
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go dd.Add(done)
	}

}
