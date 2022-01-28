package main

import "fmt"

//init函数

//执顺序执行，import的文件中包含init函数在import的时候先执行，然后在执行当前文件
//执行顺序 全局变量定义 > init函数 > main函数
var test string = testFunc()

func testFunc() string {
	fmt.Println("testFunc函数执行")
	return "1"
}

func main() {
	fmt.Println("main 函数执行")
}

//初始化函数，在main函数调用之前执行，每个源文件中都可以包含一个init函数
func init() {
	fmt.Println("init 函数执行")
}
