package main

import (
	"fmt"
	"strconv"
)

func main() {
	//类型转换只能显示转换，没有隐式转换

	//基本类型之间的转换
	var a1 int8 = 10
	var a2 float32 = float32(a1)
	fmt.Println(a1)
	fmt.Println(a2)

	//基本类型转字符串
	//方式1 fmt.Sprintf("%参数",表达式)
	var a3 string = fmt.Sprintf("%d", a1)
	fmt.Println(a3)
	fmt.Printf("%T", a3)

	var a4 float32 = 12.3
	var a5 string = fmt.Sprintf("%f", a4)
	fmt.Println(a5)
	//%q  以字符串的形式展示出来
	fmt.Printf("%T,%q", a5, a5)
	//方式2 使用strconv包的函数
	var a6 int64 = 30
	var s1 string = strconv.FormatInt(a6, 10)
	fmt.Printf("\n s1 的值 %q，s1的类型 %T", s1, s1)

	var a7 bool = true
	var s2 string = strconv.FormatBool(a7)
	fmt.Printf("\n s2 的值 %q，s2的类型 %T", s2, s2)

	//字符串转基本类型,类型无法转换时输出默认值

	var s3 string = "10"
	var s4 string = "true"
	var s5 string = "10.02"

	fmt.Println()
	var i1 int64
	i1, _ = strconv.ParseInt(s3, 10, 64)
	fmt.Println(i1)

	var b1 bool
	b1, _ = strconv.ParseBool(s4)
	fmt.Println(b1)

	var f1 float64
	f1, _ = strconv.ParseFloat(s5, 64)
	fmt.Println(f1)

}
