package main

func main() {
	// map 并发问题

	// 新建一个map
	m := make(map[int]int)

	go func() {
		for {
			_ = m[1] // 访问key=1的元素值
		}
	}()

	go func() {
		for {
			m[2] = 2 // 写入key=2的元素
		}
	}()

	select {} // 阻塞, 卡住所有协程, 比for {} 节约cpu资源, 接近休眠效果
}
