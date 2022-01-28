package main

import "fmt"

// 匿名函数

func main() {

	//匿名函数 使用场景只需要执行一次的函数
	//定义并执行一个匿名函数，将执行后的结果赋值给 result
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)

	fmt.Println(result)

	//定义一个匿名函数赋值给一个变量
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}
	//调用sub等价于执行匿名函数
	fmt.Println(sub(20, 10))
}
