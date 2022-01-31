package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//利用缓存流读取文件,适合读取大文件
	//打开文件
	file, err := os.Open("d:/aaa.txt")
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}

	//关闭文件
	defer file.Close()

	//创建一个文件流
	reader := bufio.NewReader(file)
	//读取文件内容
	for {
		content, err := reader.ReadString('\n') //读到一个换行结束
		if err == io.EOF {                      //io.EOF 已经读到了文件的结尾
			break
		}
		fmt.Println(content)
	}
	fmt.Println("文件读取完成")

	//基本文件读取，适合读取小文件  ReadFile 方法已经封装了文件的打开和关闭操作，不需要在进行文件的Open和Close操作
	content, err := ioutil.ReadFile("d:/aaa.txt")
	if err != nil {
		fmt.Println("文件读取出错", err)
		return
	}

	fmt.Println("读取到的文件内容：", string(content))
}
