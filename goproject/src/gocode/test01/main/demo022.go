package main

import (
	"fmt"
	"gocode/test01/test/structTest"
)

//挎包结构体的使用
func main() {

	student := structTest.Student{Name: "张三", Age: 15}
	fmt.Println(student)

	//挎包使用私有构造体的方式：使用私有构造体中提供的公有方法
	p1 := structTest.NewPerson("wangwu", 20)
	fmt.Println(p1)
	p1.SetName("haha")
	p1.SetAge(23)
	fmt.Println(p1.GetName())
	fmt.Println(p1.GetAge())
}
