package main

import (
	"fmt"
	//引入自定义包的文件夹路径
	//"gocode/test01/test"
	//给包起别名 在使用时也必须使用包的别名 testpa是包的别名
	testpa "gocode/test01/test"
)

//main包是程序的入口包，main函数必须放到main包下，否则无法编译
//一个目录（包）下不能有同名的函数
//函数、变量首字母大写可别其他包中使用
//同一个目录下的文件归属于一个包（包名必须相同）
//在调用别的包中的函数时要定位到函数所在的包
//使用import 引入包
//package进行包的声明
func main() {

	//调用别的包中的函数
	//test.DbCon()
	//给包起别名 在使用时也必须使用包的别名
	testpa.DbCon()
	fmt.Println("外部方法调用完成")
}
