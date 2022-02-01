package main

import "fmt"

//管道的定义及使用 创建语法  var 管道名 chan 管道类型
//是个线程安全的队列
func main() {
	//声明一个管道
	//默认创建的是可读可写的管道
	var chanNum chan int

	//创建只写的管道
	var writeOnly chan<- int
	writeOnly = make(chan int, 1)
	writeOnly <- 10
	//numW := <-readOnly //会报错
	fmt.Println(len(writeOnly))
	//创建只读的管道
	var readOnly <-chan int
	readOnly = make(chan int, 1)
	if len(readOnly) > 0 {
		fmt.Println(<-readOnly)
	}

	//初始化管道,参数1 管道的类型，参数2 管道可存放的数据量
	chanNum = make(chan int, 3)
	//向管道中写入数据
	chanNum <- 10
	chanNum <- 20
	chanNum <- 30
	//只写不读满了之后会抛出deadlock! 异常，一边放一边取满了不会抛出异常deadlock!
	//chanNum <- 40

	//从管道中读取数据
	//num1 := <-chanNum
	//num2 := <-chanNum
	//num3 := <-chanNum
	//fmt.Println(num1)
	//fmt.Println(num2)
	//fmt.Println(num3)

	//遍历管道,在遍历管道前需要先管理管道，否则管道没有数据会报deadlock，关闭了管道后没有数据会正常结束
	close(chanNum)
	for v := range chanNum {
		fmt.Println(v)
	}
}
