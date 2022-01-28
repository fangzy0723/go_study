package main

import (
	"errors"
	"fmt"
)

//异常处理

func main() {
	//
	testRecover()
	fmt.Println("testRecover 函数执行完了")
	fmt.Println("===========================")

	err := testCustomerError()
	if err != nil {
		//此时程序不会中断
		fmt.Println("函数返回了自定义异常：", err)

		//想要程序中断需要调用panic函数
		panic(err)
	}

	fmt.Println("testCustomerError 函数执行完了")
}

//defer + recover 捕获程序异常
func testRecover() {

	// defer后加上匿名函数的声明及调用
	defer func() {
		//在匿名函数内部,方法抛出异常后会在这被捕获到，没有异常时 err 为nil
		err := recover()
		if err != nil {
			fmt.Println("recover 捕获到了异常")
		}
	}()

	num1 := 10
	num2 := 2

	num3 := num1 / num2
	fmt.Println("两数相除的结果：", num3)
}

//使用errors包中的New函数声明自定义异常
func testCustomerError() error {

	num1 := 10
	num2 := 0

	if num2 == 0 {
		//抛出自定义异常
		return errors.New("被除数不能为0")
	} else {
		//程序往下执行
		num3 := num1 / num2
		fmt.Println("两数相除的结果：", num3)

		//程序没有异常返回一个nil
		return nil
	}
}
