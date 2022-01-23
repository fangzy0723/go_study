package main

import (
	"fmt"
	//引入自定义的包是从 $GOPATH/src/后开始算的，使用/进行路径分割,引入其他包路径到包名即可不用到文件名，使用时是 包名.属性名
	"gocode/test01/test"
)

func main() {
	//使用时是 包名.属性名
	fmt.Println(test.StudyName)
	//首字母小写包中私有不能在别的包中使用
	//fmt.Println(test.studyAge)
}
