package main

//反射  结构体

import (
	"fmt"
	"reflect"
)

// TestStudent 定义一个结构体
type TestStudent struct {
	Name string
	Age  int
}

func (t TestStudent) PrintStudent() {
	fmt.Println("学生信息：", t.Name, t.Age)
}
func (t TestStudent) SPrintStudent(age int) {
	fmt.Println("传参的值：", age)
}

func testReflectStudent(i interface{}) {
	reType := reflect.TypeOf(i)

	reValue := reflect.ValueOf(i)

	fmt.Println(reType)
	fmt.Println(reValue)

	i1 := reValue.Interface()
	testStudent := i1.(TestStudent)
	fmt.Println(testStudent.Name)

	//不需要传参的方法 给nil即可
	//方法的顺序是按照ASCII的顺序排列的
	reValue.Method(0).Call(nil)

	//需要传参的 先定义一个reflect.Value类型的切片
	var param []reflect.Value
	param = append(param, reflect.ValueOf(10))
	reValue.Method(1).Call(param)
}

func main() {

	student := TestStudent{
		Name: "张三",
		Age:  20,
	}
	testReflectStudent(student)
}
