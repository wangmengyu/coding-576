package main

import (
	"fmt"
	"time"
)

type Person struct {
	salary int // 工资
	level  int // 等级
}

// promote 升值加薪
func (p *Person) promote() {

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
	for i := 1; i <= 200; i++ {
		go p.promote()
	}

	time.Sleep(10 * time.Second)

}
