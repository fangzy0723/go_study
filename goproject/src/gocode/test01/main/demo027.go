package main

import (
	"bufio"
	"fmt"
	"os"
)

//文件写入

func main() {

	//O_CREATE 文件不存在是创建
	//perm:文件权限
	file, err := os.OpenFile("d:/aaa.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	defer file.Close()
	//创建一个写流
	write := bufio.NewWriter(file)
	//向文件写入内容
	for i := 0; i < 10; i++ {
		write.WriteString("你好 \n")
	}
	//清空系统缓冲区，刷到磁盘
	write.Flush()

	//输出0666的文件模式，适用于linux系统
	str := os.FileMode(0666).String()
	fmt.Println(str)
}
