package main

import "fmt"

//函数的声明
/*
参数可声明为可变参数 如  args... int
没有返回值时 返回值类型可不写
只有一个参数可不加()

函数命名遵循驼峰形式
函数名首字母大写可在其他包中使用 类似于java中的public
函数名首字母小写只能在本包中使用 类似于java中的private

格式：
func 函数名 (参数名 参数类型,参数名 参数类型) (返回值类型,返回值类型){

	return value
}

基本类型和数组的参数都是值传递,函数内的参数变化不会影响原值；函数内改变原值可使用可使用指针进行地址传递

*/

//计算两个参数的和 , 只有一个返回值时 返回值类型的()可不写
func f1(num1 int, num2 int) int {
	return num1 + num2
}

//计算两个参数的和并输出
func f2(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("f2.sum:", sum)
}

//计算两个参数的和、差
func f3(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	cha := num1 - num2
	return sum, cha
}

//可变参数,输出可变参数的值
func f4(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

//使用指针进行地址传递，在函数内改变参数的值
func f5(num *int) {
	*num = 30
}

//接收一个函数类型的参数，函数参数名：fv 函数类型： func(int2 int, num int) 函数返回值类型： int
func f6(fv func(int2 int, num int) int) {
	fmt.Println("f6开始执行")
	//
	v := fv(1, 2)
	fmt.Println("f6内部调用函数类型的参数计算的结果：", v)
}

//计算两个参数的和、差 函数返回值命名
func f7(num1 int, num2 int) (sum int, cha int) {
	//顺序可以不用对应
	sum = num1 + num2
	cha = num1 - num2
	return
}

type myFuncType func(int2 int, num int) int

//接收一个函数类型的参数，使用自定义函数类型
func f8(fv myFuncType) {
	fmt.Println("f8开始执行")
	//
	v := fv(1, 2)
	fmt.Println("f8内部调用函数类型的参数计算的结果：", v)
}

func main() {
	sum := f1(10, 20)
	fmt.Println("sum:", sum)
	f2(10, 12)
	//sum1, _ := f3(10, 20)  可使用_忽略不需要关注的返回值
	sum1, cha := f3(10, 20)
	fmt.Printf("sum1:%v  cha: %v", sum1, cha)
	f4(1, 2, 3, 4, 5, 6, 7, 8, 9)
	//在函数内改变num的值
	var num int = 10
	f5(&num) //调用参数传递num的内存地址
	fmt.Println("num 的值为：", num)

	//函数赋值为一个变量
	v1 := f1
	fmt.Printf("v1的类型 %T  f1的类型 %T \n", v1, f1)
	v1v := v1(1, 2) //等价f1(1, 2)
	f1v := f1(1, 2)
	fmt.Println("v1v:", v1v)
	fmt.Println("f1v:", f1v)

	//把函数f1 当参数传递给f6
	f6(f1)
	//函数命名
	sun1, cha1 := f7(10, 20)
	fmt.Println(sun1, cha1)

	//自定义类型  相当于起别名
	//语法  type 自定义数据类型命 数据类型
	// 给int类型起个别名为myInt ，但是 int 和myInt 不是同一种类型
	type myInt int
	var num2 myInt = 20
	fmt.Println(num2)

	//函数参数类型使用自定义函数类型
	f8(f1)
}
