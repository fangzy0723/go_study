package main

import "fmt"

//结构体的继承

// A 定义一个A的结构体
type A struct {
	name string
	age  int
}

func (a A) eat() {
	fmt.Println("调用了A的吃饭方法")
}

func (a A) paXing() {
	fmt.Println("调用了A的爬行方法")
}

// C 定义一个C的结构体
type C struct {
	name string
}

func (c *C) paXing() {
	fmt.Println("调用了C的爬行方法")
}

// B 定义一个B的结构体，继承A的结构体
type B struct {
	A       //继承匿名结构体
	c    *C //属性的类型是个结构体指针
	name string
}

func (b *B) eat() {
	fmt.Println("调用了B的吃饭方法")
}

// 结构体中的属性和继承的结构体中有同名的属性，就近使用
// 多继承时需要,包含同名方法需要指定使用哪个结构体的方法
// 属性的类型是结构体需要使用 属性名.结构体属性名
func main() {

	b := B{A: A{name: "A", age: 10}, c: &C{name: "C"}, name: "B"}
	fmt.Println(b)
	fmt.Println(b.name) //就近使用 输出B
	fmt.Println(b.age)  //等价于 b.A.age
	fmt.Println(*(b.c))
	b.c.paXing()
	b.eat()
	b.A.eat()
	b.A.paXing()

}
