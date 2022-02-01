package main

import (
	"fmt"
	"sync"
	"time"
)

//启动协程的方式   加go关键字
//特点：主死从随  主线程结束，主线程里面的协程一并结束不管有没有属性完

//主线程等待所有协程执行完在结束，类似于java中的CountDownLatch
var wg sync.WaitGroup

func testXieCheng() {
	fmt.Println("使用协程执行test函数")
	wg.Done() //
}

func main() {

	wg.Add(2) //启动两个协程
	//启动一个协程执行test函数
	go testXieCheng()

	go func() {
		fmt.Println("在匿名函数内部执行")
		time.Sleep(time.Second) //休眠1秒
		wg.Done()               // 协程执行结束调用down方法
	}()
	wg.Wait() //等待指定个数协程执行结束
	fmt.Println("testXieCheng函数执行结束")
}
