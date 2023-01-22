package main

import "sync"

func main() {

	// 锁复制引起的问题
	m := sync.Mutex{}
	m.Lock()
	n := m
	m.Unlock()
	n.Lock()
	n.Unlock()

}
