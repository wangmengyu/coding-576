package main

import "fmt"

// ICar  定义car的接口
type ICar interface {
	Drive()
}

type Trunk struct {
	// 卡车
	Model string
}
type ITrafficTool interface {
	Drive()
}

func (t *Trunk) Drive() { // 实现car的drive接口.
	fmt.Println(t.Model)
}

func main() {
	// 定义一个ICAR的接口变量
	var c ICar = &Trunk{Model: "123"}
	fmt.Println("%T", c) // 双击shift查询iface查看, 选择Type类型, 可看到地一个就是
	/**
	type iface struct {
		tab  *itab //
		data unsafe.Pointer // 万能指针, 指向数据, 本例是指向了Trunk类型.
	}
	type itab struct {
		inter *interfacetype // 指向的接口的类型.
		_type *_type // 当前装载的值具体是什么类型
		hash  uint32 // copy of _type.hash. Used for type switches.
		_     [4]byte
		fun   [1]uintptr // 函数, 指明了当前type实现了哪些方法, variable sized. fun[0]==0 means _type does not implement inter.
	}
	*/
	// 本质上c 是一个 iface 类型的结构提.
	// 断言, 类型转换成具体的类
	tk := c.(*Trunk)
	tk.Drive()

	tt := c.(ITrafficTool) // 可以从一种接口转换成另一种接口
	fmt.Printf("%+v, %T\n", tt, tt)
	switch c.(type) {
	case ICar:
		fmt.Println("id ICar")
	case ITrafficTool:
		fmt.Println("is ITrafficTool")
	default:

	}
}
