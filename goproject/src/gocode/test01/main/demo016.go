package main

import (
	"fmt"
	"strings"
)

//字符串函数的使用
//更多详细用法可参考：https://studygolang.com/pkgdoc
func main() {

	var str1 string = "A"
	var str2 string = "a"
	fmt.Println("不区分大小写字符串比较", strings.EqualFold(str1, str2))
	fmt.Println("区分大小写字符串比较", str1 == str2)

	//判断字符串中是否包含子字符串
	b1 := strings.Contains("北京天安门", "北京")
	b2 := strings.Contains("北京天安门", "北大")
	fmt.Println(b1, b2)
}
