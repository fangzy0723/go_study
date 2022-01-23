package main

import "fmt"

//运算符示例
func main() {
	//+   可表示 1、正数  2、两个整数相加  3、 两个字符串拼接
	var n1 int = +10
	var n2 int = 10 + 20
	var n3 string = "hello" + "hello"

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println("================")

	// /除号
	fmt.Println(10 / 3)   //都是整数参数运算的结果也是整数
	fmt.Println(10.0 / 3) //有浮点数参数运算的结果是浮点数
	fmt.Println("================")

	//取模，运算结果正负只看 前面数字的符号
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)
	//自增
	var n4 int = 10
	n4++
	fmt.Println(n4)

	//++  -- 只能单独使用 不能参与到运算中，只能写到变量名的后面不能写到变量名的前面
	//自减
	n4--
	fmt.Println(n4)
	fmt.Println("=================")

	//赋值运算
	var n5 int = 10
	n5 += 10
	fmt.Println(n5)

	n5 -= 5
	fmt.Println(n5)

	n5 *= 2
	fmt.Println(n5)

	n5 /= 10
	fmt.Println(n5)

	fmt.Println("=================")
	//比较运算
	fmt.Println(10 < 20)
	fmt.Println(10 == 20)
	fmt.Println(10 > 20)
	fmt.Println(10 <= 20)
	fmt.Println(10 >= 20)

	fmt.Println("=================")
	//逻辑运算
	//与运算  全为真才为真
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false && false)

	//或运算  有为真就为真
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false || false)

	//非运算  取反
	fmt.Println(!true)

	//其他运算  & 取一个变量的地址值  * 获取指针指向的具体值

	var n6 int = 10
	var n7 *int = &n6
	fmt.Println(n7)
	fmt.Println(*n7)

	fmt.Println("==================")
	//接收控制台输入的参数，接收用户控制台输入参数需要使用变量的地址值
	//方式1
	var age int
	fmt.Println("请输入用户的年龄：")
	fmt.Scanln(&age)
	fmt.Println("用户输入的年龄为:", age)

	//方式2
	var name string
	fmt.Println("请输入用户的姓名：")
	fmt.Scanf("%s", &name)
	fmt.Println("用户输入的姓名是：", name)

}
