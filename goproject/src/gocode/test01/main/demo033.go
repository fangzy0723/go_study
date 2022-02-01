package main

import (
	"fmt"
	"sync"
	"time"
)

//定义两个协程操作同一个管道，一个协程向管道内写入数据，另一各协程向从管道读取数据

var wgc sync.WaitGroup

//写函数 向管道内写入数据
func writeChan(ch chan int) {
	defer wgc.Done()
	defer close(ch) //写完之后不关闭管道 读取结束会报错
	for i := 0; i < 20; i++ {
		ch <- i
		fmt.Println("向管道内写入数据：", i)
		time.Sleep(time.Second)
	}
}

//读函数 从管道内读取数据
func readChan(ch chan int) {
	defer wgc.Done()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("从管道内读取数据出错")
		}
	}()

	for v := range ch {
		fmt.Println("从管道内读取到的内容：", v)
		time.Sleep(time.Second * 1)
	}

}

func main() {
	wgc.Add(2)
	//初始化一个管道
	var intChan chan int = make(chan int, 10)

	//定义两个协程，分别向管道内写入数据和取出数据
	go writeChan(intChan)
	go readChan(intChan)
	wgc.Wait()
	fmt.Println("协程执行结束")
}
