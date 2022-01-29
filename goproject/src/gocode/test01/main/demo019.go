package main

import (
	"fmt"
)

//切片的定义及使用
//切片是有3个字段的数据结构，1、指向底层数组的指针  2、切片的长度  3、切片的容量
func main() {

	//切片的定义
	//方式1  基于已有数组创建切片
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	//arr[1,4]  截取数组arr下标为[1-4）的元素给slice1
	//slice1和数组arr[1]指向相同的内存地址
	//var slice1 []int = arr[1:4]
	slice1 := arr[1:4] //上面一行代码的简写形式
	fmt.Printf("数组arr下标为1的内存发地址:%p \n", &arr[1])
	fmt.Printf("切片slice1的内存地址：%p \n", slice1)
	fmt.Printf("切片slice1[0]的内存地址：%p \n", &slice1[0])
	fmt.Printf("切片slice1的长度：%v 容量：%v \n", len(slice1), cap(slice1))
	//修改切片中元素的值也会一起改原数组的元素值
	slice1[1] = 10
	fmt.Println(arr)
	fmt.Println(slice1)

	//方式2 使用mark函数创建切片，3个参数分别为 切片的类型，切片的长度，切片的容量
	//会在底层创建一个数组，该数组对外不可见，不能直接操作这个数组
	slice2 := make([]int, 3, 8)
	fmt.Println(slice2)

	//方式3  定义一个切片直接指向具体的数组
	slice3 := [3]int{2, 3, 4}
	fmt.Println(slice3)

	//append 给切片追加元素，返回一个新的切片
	//对老数组进行扩容时生成新的数组，将老数组的值复制到新数组中，再将新元素追加到新数组中
	slice2 = append(slice2, 5, 6, 7)
	fmt.Println(slice2)

	//切片拷贝
	//将slice1中的元素拷贝到slice2中,对应下标的元素进行拷贝，目标数组的超出原数组的部分不做任何操作
	copy(slice2, slice1)
	fmt.Println(slice2)

	slice4 := make([]int, 2, 3)
	copy(slice4, slice2)
	fmt.Println(slice4)

	//切片定义后不能直接使用需要让其引用一个数组
	//不能越界使用
	//切片可以继续切片 开始的0可不写  结束位置是数组的结尾可不写
	//slice1 := arr[0:end]  -- >> slice1 := arr[:end]
	//slice1 := arr[1:len(arr)]  -- >> slice1 := arr[1:]
	//slice1 := arr[0:len(arr)]  -- >> slice1 := arr[:]

	//切片的遍历 方式同数组的遍历
	//方式1
	for i := 0; i < len(slice4); i++ {
		fmt.Println(slice4[i])
	}

	//方式2
	for i, v := range slice4 {
		fmt.Printf("下标 %v 元素 %v \n", i, v)
	}

}
