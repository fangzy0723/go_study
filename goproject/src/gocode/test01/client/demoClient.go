package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	//客户端启动

	fmt.Println("客户端启动。。。")
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("客户端连接服务端失败")
		return
	}

	fmt.Println("客户端连接服务端成功")
	defer conn.Close()
	for {
		//os.Stdin 从终端输入
		read := bufio.NewReader(os.Stdin)
		content, err01 := read.ReadString('\n')
		if err01 != nil {
			fmt.Println("从终端获取输入失败", err01)
			return
		}
		//发送内容给服务端
		n, err02 := conn.Write([]byte(content))
		if err02 != nil {
			fmt.Println("给服务端发送内容失败")
			return
		}
		fmt.Printf("接收终端输入的内容发送给服务端 共%d个字节 \n", n)

		for {
			buf := make([]byte, 1024)
			n, err03 := conn.Read(buf)
			if err03 != nil {
				fmt.Println("接收服务端发送内失败", err03)
				return
			}
			fmt.Println("服务端返回内容：", string(buf[0:n]))
			if n < 1024 {
				break
			}
		}
	}

}
