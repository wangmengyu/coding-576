package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu     sync.Mutex // 要实现互斥锁, 需要再结构体中添加sync.Mutex成员.
	salary int        // 工资
	level  int        // 等级
}

// promote 升值加薪
func (p *Person) promote() {
	p.mu.Lock()         // 操作之前加锁
	defer p.mu.Unlock() // 习惯在lock之后加defer的解锁语句. 防止过程中出现panic.
	p.salary = p.salary + 1000
	p.level = p.level + 1

	fmt.Println(p.salary)
	fmt.Println(p.level)
	//p.mu.Unlock() //  这里就不需要Unlock 了, 因为defer 会解锁.
}

func main() {
	p := Person{
		salary: 10000,
		level:  1,
	}

	// 需求: 在前一个升职加薪后再能执行后一个升职加薪操作:
	go p.promote()
	go p.promote()
	go p.promote()

	time.Sleep(1 * time.Second)

}
