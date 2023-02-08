package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	mu     sync.RWMutex // 结构体中定义 读写锁
	salary int          // 工资
	level  int          // 等级
}

// promote 升值加薪
func (p *Person) promote() { // 写操作中使用 Lock 和 Unlock 来调用写锁的使用
	p.mu.Lock()         // 操作之前加锁
	defer p.mu.Unlock() // 习惯在lock之后加defer的解锁语句. 防止过程中出现panic.
	p.salary = p.salary + 1000
	p.level = p.level + 1

	fmt.Println(p.salary)
	fmt.Println(p.level)
	//p.mu.Unlock() //  这里就不需要Unlock 了, 因为defer 会解锁.
}

// 打印个人信息, [纯读取]
func (p *Person) printPerson() { // 在读操作中使用 RLock 和 RUnlock 来控制读锁的使用
	p.mu.RLock()         // 读取锁
	defer p.mu.RUnlock() //读取解锁
	fmt.Println(p.salary)
	fmt.Println(p.level)
}

func main() {
	p := Person{
		salary: 10000,
		level:  1,
	}

	go p.printPerson() //  并发读. 读取锁
	go p.printPerson() //  并发读. 读取锁
	go p.printPerson() //  并发读 . 读取锁
	go p.promote()     // 会等前面三个都做完才做, 用了互斥锁
	time.Sleep(1 * time.Second)

}
