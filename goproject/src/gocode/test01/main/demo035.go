package main

//反射  基本类型

import (
	"fmt"
	"reflect"
)

//利用一个函数  函数的参数定义为一个空接口
func testReflect(i interface{}) {

	//获取类型 reflect.Type
	reType := reflect.TypeOf(i)
	fmt.Println(reType)
	fmt.Printf("reType的具体类型 %T \n", reType)

	//获取value    reflect.Value
	reValue := reflect.ValueOf(i)

	fmt.Println(reValue)
	fmt.Printf("reValue的具体类型 %T", reValue)

	//value因为不是int类型不能直接使用，需要转换
	num := 20 + reValue.Int()
	fmt.Println("nunm:", num)

	//已知reflect.Value，反向转换为该值的类型
	i1 := reValue.Interface()
	n1 := i1.(int)
	fmt.Println("n1:", n1)
}

//通过反射修改字段的值
func setValue(i interface{}) {
	reValue := reflect.ValueOf(i)
	reValue.Elem().SetInt(30)
}

func main() {

	num := 20
	testReflect(num)

	setValue(&num) //默认是值传递，修改值需要传入指针
	fmt.Println("修改之后num的值：", num)

}
