package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println("当前系统时间：", now)

	year := now.Year()
	month := now.Month()
	m := int(now.Month())
	fmt.Println(year, month, m)

	//将时间字符串2022-01-27 10:20:12转成Time类型
	time, _ := time.Parse("2006-01-02 15:04:05", "2022-01-27 10:20:12")
	//now 是否在time之后
	fmt.Println(now.After(time))
	//now 是否在time之前
	fmt.Println(now.Before(time))

	//对当前时间进行格式化成时间字符串
	//2006、01、02 15、04、05 分别对应的是年月日 时分秒，固定写法
	nowStr := now.Format("2006-01-02 15:04:05")
	fmt.Println("格式换后的当前系统时间：", nowStr)
	//输出本地时区的时间
	fmt.Println(now.Local())
}
