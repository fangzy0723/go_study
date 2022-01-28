package main

import "fmt"

// 闭包
//定义一个函数 函数的返回值是一个函数
func getSum() func(num int) int {

	//变量sum的值在内存中一直存在，保存的是修改之后的值
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {

	f := getSum()
	fmt.Println(f(1)) //1
	fmt.Println(f(2)) //2
	fmt.Println(f(3)) //6
}
