package main

import (
	"fmt"
	"sync"
)

type Person struct {
	mu     sync.RWMutex // 读写锁
	salary int          // 工资
	level  int          // 等级
}

// promote 升值加薪
func (p *Person) promote(wg *sync.WaitGroup) {
	p.mu.Lock()         // 操作之前加锁
	defer p.mu.Unlock() // 习惯在lock之后加defer的解锁语句. 防止过程中出现panic.
	p.salary = p.salary + 1000
	p.level = p.level + 1

	fmt.Println(p.salary)
	fmt.Println(p.level)
	//p.mu.Unlock() //  这里就不需要Unlock 了, 因为defer 会解锁.
	wg.Done()
}

func main() {
	p := Person{
		salary: 10000,
		level:  1,
	}
	wg := sync.WaitGroup{}
	wg.Add(3)         // 开3个协程
	go p.promote(&wg) // 会等前面三个都做完才做, 用了互斥锁
	go p.promote(&wg) // 会等前面三个都做完才做, 用了互斥锁
	go p.promote(&wg) // 会等前面三个都做完才做, 用了互斥锁
	// 都完成时退出
	wg.Wait()
	fmt.Println("done")

}
