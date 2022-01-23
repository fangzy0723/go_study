package main

//变量的声明及使用
import "fmt"

//声明全局变量
var d3 = 10
var d4 = 12.4

//同时声明多个全局变量
var (
	d5 = "haha"
	d6 = 10
)

func main() {
	//声明局部变量
	//方式1 先声明变量后复制

	var age int
	age = 10
	fmt.Println("age=", age)
	//方式2 声明变量并赋值
	var name = "lisa"
	fmt.Println("name=", name)

	//方式3 声明变量没有赋值  使用初始值
	var num int
	fmt.Println("num", num)

	//方式4 省略var 关键词  使用:=  注意不能使用=

	addr := "北京市"
	fmt.Println("地址", addr)

	//、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、
	//同时声明多个变量

	var a1, a2, a3 int = 10, 20, 30
	fmt.Println("a1=", a1, "a2=", a2, "a3=", a3)

	var s1, s2, s3 = 10, "张三", 10.2
	fmt.Println("s1=", s1, "s2=", s2, "s3=", s3)

	d1, d2 := 10, "lisa"
	fmt.Println("d1=", d1, "d2=", d2)

	fmt.Println(d3)
	fmt.Println(d4)
	fmt.Println(d5)
	fmt.Println(d6)
}
