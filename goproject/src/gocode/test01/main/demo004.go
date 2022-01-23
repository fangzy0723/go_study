package main

import "fmt"

//指针变量的声明及使用
func main() {
	//获取一个变量的内存地址
	var age int = 20
	fmt.Println("age变量的内存地址", &age)

	//声明一个指针类型的变量，指针就是内存地址

	//ptr 指针变量名，类型是 *int（指向int类型的指针）
	//ptr 的值是一个内存地址，获取对应内存地址的数据值使用*号
	var ptr *int = &age
	fmt.Println("ptr的值", ptr, "内存地址的数据值", *ptr)

	//通过指针可以改变变量的值
	*ptr = 25
	fmt.Println(age)

	//指针接收的是一个内存地址

}
