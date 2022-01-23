package main

import "fmt"

func main() {

	//分支判断 表达式的() 可不写
	if count := 10; count < 20 {
		fmt.Println("库存不足")
	} else {
		fmt.Println("库存充足")
	}

	//多分支判断
	var score int8 = 55
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	var num int = 1
	switch num {
	case 10: //case 后面可以写多个值 , 隔开
		fmt.Println("10")
		//不需要break
	case 9:
		fmt.Println("9")
	case 8:
		fmt.Println("8")
	default:
		//兜底。以上都不满足时执行
		fmt.Println("0")
	}
}
