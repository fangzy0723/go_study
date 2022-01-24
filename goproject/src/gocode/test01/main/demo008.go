package main

import "fmt"

// for 循环的使用  关键字  break（结束离他最近的for循环） continue（结束本次循环开始下一次循环） return（结束方法的执行）  goto（跳转到某一行去执行）
func main() {

	//输出0-10的数字
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//for range 的使用,rang 对 str字符串进行遍历  下标给i，值给value
	var str string = "hello go 中国"
	for i, value := range str {
		fmt.Printf("下标：%v 值：%c \n", i, value)
	}

	//计算0-100的和 当和大于等于100 时退出for循环
	var sum int = 0
	for i := 0; i <= 100; i++ {
		sum += i
		if sum >= 300 {
			break //退出离它最近的for循环,break 在switch中使用可不写

			//return //退出当前函数的执行
		}
	}
	fmt.Println("sum:", sum)

	for i := 0; i < 5; i++ {
		for j := 1; j <= 3; j++ {

			if i == 2 && j == 2 {
				//continue //结束本次循环 开始下一次循环
				break //结束最近的for循环
			}
			fmt.Printf("i:%v j:%v \n", i, j)
		}
	}

lable1:
	for i := 0; i < 5; i++ {
		for j := 1; j <= 3; j++ {

			if i == 2 && j == 2 {
				//continue //结束本次循环 开始下一次循环
				continue lable1 //结束指定标签的循环 开始下一次循环
				//break //结束最近的for循环
				//break lable1 //结束标签标识的for循环
			}
			fmt.Printf("i:%v j:%v \n", i, j)
		}
	}

	fmt.Println("m1")
	fmt.Println("m2")
	goto leable2 // 当执行到这行是跳转到 leable2 标签标识的位置处执行
	fmt.Println("m3")
	fmt.Println("m4")
	fmt.Println("m5")
leable2:
	fmt.Println("m6")
	fmt.Println("m7")

}
