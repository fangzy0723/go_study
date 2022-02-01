package main

//select 功能：解决多个管道的选择问题，多路复用，可以从多个管道中随机公平的选择一个来执行

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int, 1)
	strCh := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 3)
		intCh <- 10
	}()

	go func() {
		time.Sleep(time.Second * 2)
		strCh <- "北京"
	}()

	//time.Sleep(time.Second * 3)

	select {
	case v := <-intCh:
		fmt.Println(v)
	case v := <-strCh:
		fmt.Println(v)
		//default: //防止阻塞
		//	fmt.Println("未获取到数据")
	}
}
