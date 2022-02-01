package main

//互斥锁的使用示例

import (
	"fmt"
	"sync"
)

var totalNUm int64 = 0

//目的是协程全部执行完主线程在往下走
var wg1 sync.WaitGroup

//互斥锁，读写加同一把锁实现读写互斥，同时只有一个协程操作数据
var lock sync.Mutex

func add() {

	for i := 0; i < 10000; i++ {
		lock.Lock() //加锁
		totalNUm += 1
		lock.Unlock() //释放锁
	}
	wg1.Done() // 或者使用defer 关键字
}

func sub() {
	for i := 0; i < 10000; i++ {
		lock.Lock() //加锁
		totalNUm -= 1
		lock.Unlock() //释放锁
	}
	wg1.Done()
}

func main() {
	wg1.Add(2) //数量是需要创建的协程的数量
	go add()
	go sub()
	wg1.Wait()
	fmt.Println("totalNum", totalNUm)
}
