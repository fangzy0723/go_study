package main

//一个类实现多个接口，需要实现这些接口中的所有方法

import "fmt"

//定义接口A1
type A1 interface {
	a1()
}

//定义接口B1
type B1 interface {
	b1()
}

//定义一个结构体实现A1、B1接口
type C1 struct {
}

func (c C1) a1() {
	fmt.Println("这是实现A1接口中的方法")
}
func (c C1) b1() {
	fmt.Println("这是实现B1接口中的方法")
}

func main() {
	var c1 C1 = C1{}
	var a1 A1 = c1
	var b1 B1 = c1

	//可以使用接口调用
	a1.a1()
	b1.b1()
	//也可以使用接口实现类的实例调用
	c1.a1()
	c1.b1()

}
