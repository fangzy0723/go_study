package main

//读写锁  适用场景 读多写少

import (
	"fmt"
	"sync"
	"time"
)

var sumNum int

var wg2 sync.WaitGroup

//声明一个读写锁
var rwLock sync.RWMutex

//读方法
func read() {
	rwLock.RLock() //加读锁  在获取锁的时候不影响别的协程获取这个读锁
	fmt.Println("读方法开始")
	time.Sleep(time.Second)
	fmt.Println("读方法结束")
	rwLock.RUnlock() //释放读锁

	wg2.Done()
}

//写方法
func write() {
	rwLock.Lock() //加写锁  在获取写锁的时候别的协程不能获取读锁和写锁
	fmt.Println("写方法开始")
	time.Sleep(time.Second * 5)
	fmt.Println("写方法结束")
	rwLock.Unlock() //释放写锁

	wg2.Done()
}
func main() {

	wg2.Add(6)
	//模仿对多写少的的场景
	for i := 0; i < 5; i++ {
		go read()
	}

	go write()

	wg2.Wait()
}
