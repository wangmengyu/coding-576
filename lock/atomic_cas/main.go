package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Person struct {
	mu     int32 // 使用atomic锁
	salary int   // 工资
	level  int   // 等级
}

// promote 升值加薪
func (p *Person) promote() {
	atomic.CompareAndSwapInt32(&p.mu, 0, 1) // 操作之前加锁 , 如果=0进行+1
	p.salary = p.salary + 1000
	p.level = p.level + 1

	fmt.Println(p.salary)
	fmt.Println(p.level)
	atomic.CompareAndSwapInt32(&p.mu, 1, 0) //  操作完后解锁
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
