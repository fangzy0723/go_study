package main

import "fmt"

// 一个接口多个实现类
//一个结构体要实现一个接口。必须要实现这个接口中的所有方法，没有implement关键字不需要显示的实现接口，只要实现了接口中的所有方法就认为是实现了这个接口
//接口类型默认是指针类型，使用时需要先初始化，接口本身不能创建示例，需要指向一个实现了该接口的自定义类型的实例
//一个接口继承了另外一个接口，在实现的时候需要实现这个继承的接口中的方法

// AI A 订单一个A接口
type AI interface {
	a()
}

// BS B 定义一个结构体B 实现AI接口中的方法
type BS struct {
}

func (b BS) a() {
	fmt.Println("这是B结构体实现A接口的方法")
}
func (b BS) bf() {
	fmt.Println("这是BS中特有的方法")
}

// CS 定义一个C接口实现AI中的方法
type CS struct {
}

func (c CS) a() {
	fmt.Println("这是C结构体实现A接口的方法")
}
func (c CS) cf() {
	fmt.Println("这是CS中特有的方法")
}

//定义一个函数 根据不对的对象调用不同的实现方法
// 多态的形式
func testInter(a AI) {
	a.a()
	//通过断言判断该调用实现类中哪个特有的方法
	//bs,flag := a.(BS)//flag:true 说明a是BS类型的实例，false说明不是
	switch a.(type) { //type是go中的一个关键字，固定写法
	case BS:
		bs := a.(BS)
		bs.bf()
	case CS:
		cs := a.(CS)
		cs.cf()
	}
}

func main() {
	var ai1 AI = BS{}
	var ai2 AI = CS{}

	testInter(ai1)
	testInter(ai2)
}
