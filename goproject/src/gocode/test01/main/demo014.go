package main

import "fmt"

//defer的使用场景，在函数执行完之后释放资源
//特点：遇到defer关键字，会将后面的代码语句压入栈，同时将相关的值一并拷贝到粘中，不受后续对变量中的改变而改变
//函数在执行完毕之后再从栈中拿出需要语句进行执行(栈是先进后出，后面的代码会先执行)

func testDefer(num1 int, num2 int) int {

	defer fmt.Println("num1:", num1)
	defer fmt.Println("num2:", num2)
	sum := num1 + num2
	fmt.Println("testDefer 函数执行完毕")
	return sum
}

func main() {

	fmt.Println(testDefer(10, 20))
}
