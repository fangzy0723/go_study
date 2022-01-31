package main

import (
	"fmt"
	"io/ioutil"
)

//读一个文件内容并写入到另外一个文件

func main() {
	//读文件名
	readFileName := "d:/aaa.txt"
	//写文件名
	writeFileName := "d:/bbb.txt"

	readFileContent, err := ioutil.ReadFile(readFileName)
	if err != nil {
		fmt.Println("读取文件内容失败", err)
		return
	}

	err1 := ioutil.WriteFile(writeFileName, readFileContent, 0666)
	if err1 != nil {
		fmt.Println("文件写入出错", err1)
		return
	}

}
