package main

//导包
import (
	"fmt"
	"unsafe"
)

//整数类型的定义
func main() {
	//有符号整数
	//int8 取值范围   -2^7（-128） - 2^7-1(127)  1字节
	//int16 取值范围   -2^15 - 2^15-1   2字节
	//int32 取值范围   -2^31 - 2^31-1   4字节
	//int64 取值范围   -2^63 - 2^63-1   8字节

	var a1 int8 = 1
	var a2 int16 = 2
	var a3 int32 = 3
	var a4 int64 = 4

	//无符号整数
	//uint8 取值范围   0 - 2^8-1(255)   1字节
	//uint16 取值范围   0 - 2^16-1    2字节
	//uint32 取值范围   0 - 2^32-1    4字节
	//uint64 取值范围   0 - 2^64-1    8字节

	var a5 uint8 = 5
	var a6 uint16 = 6
	var a7 uint32 = 7
	var a8 uint64 = 8

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
	fmt.Println(a6)
	fmt.Println(a7)
	fmt.Println(a8)
	fmt.Printf("a1 的类型 %T", a1)
	fmt.Println()
	fmt.Println("a1 占用的字节", unsafe.Sizeof(a1))
	fmt.Println("==================================")

	//浮点型 可以表示正浮点数也可以表示负浮点数  可以使用科学计数法表示  E不区分大小写
	//float32  4字节  可能存在精度损失
	//float64  8字节
	var f1 float32 = 10.3
	var f2 float64 = 20.3
	fmt.Println(f1)
	fmt.Println(f2)

	///////////////////////////////////////////////////////
	//bool 类型的定义及使用
	var b1 bool = true
	var b2 bool = 5 > 9

	fmt.Println(b1)
	fmt.Println(b2)

	//字节的定义及使用
	var by1 byte = 1
	var by2 byte = 'A'
	var by3 byte = '('
	fmt.Println(by1)
	fmt.Println(by2)
	fmt.Println(by3)

	//字符串的定义及使用  字符串一旦定义 其值不能改变
	var name string = "lisa"
	fmt.Println(name)

	//给变量名重新赋值
	name = "zhangsan"
	fmt.Println(name)

	//定义一个多行字符串
	var desc string = `
		这是一个
多行字符串
哈哈 
你好

	`
	fmt.Println(desc)

	//字符串拼接
	name += "北京" + "中国"
	fmt.Println(name)

}
