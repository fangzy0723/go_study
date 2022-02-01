package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//获取客户端发送的内容
func getContent(conn net.Conn) {
	//用完之后关闭连接
	//defer conn.Close()

	for {
		//创建一个切片，将接收到的内容放入切片
		arr := make([]byte, 1024)
		//从连接中获取内容放入切片 n:接收到数据的字节数
		n, err := conn.Read(arr)
		if err != nil {
			fmt.Println("获取客户端发送内容失败")
			return
		}
		//将获取到的内容输入到控制台
		fmt.Println(string(arr[0:n]))
		if n < 1024 {
			break
		}
	}

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
		fmt.Println("给客户端端发送内容失败")
		return
	}
	fmt.Printf("接收终端输入的内容发送给客户端 共%d个字节 \n", n)
}

func main() {

	fmt.Println("服务端启动成功")
	listener, err := net.Listen("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("服务端启动失败", err)
		return
	}
	for {
		//等待客户端连接
		conn, err01 := listener.Accept()
		if err01 != nil {
			fmt.Println("等待客户端连接失败", err01)
		} else {
			fmt.Printf("客户端连接成功 客户端地址信息 %v \n", conn.RemoteAddr().String())
			//连接成功后创建一个协程获取客户端发送的内容
			go getContent(conn)
		}
	}

}
