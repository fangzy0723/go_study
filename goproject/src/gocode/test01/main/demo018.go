package main

import "fmt"

//数组的定义及使用
//数组每个空间占用的字节数取决于数组元素类型
//数组默认是值传递，方法间传递会进行值拷贝

func main() {

	//数组的定义
	//方式1
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)
	fmt.Printf("数组的地址 %p \n", &arr1)
	fmt.Printf("数组下标0的地址 %p \n", &arr1[0])
	fmt.Printf("数组下标1的地址 %p \n", &arr1[1])
	fmt.Printf("数组下标2的地址 %p \n", &arr1[2])
	fmt.Printf("数组的类型 %T \n", arr1) // 长度是数组类型的一部分
	//方式2
	var arr2 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr2)
	//方式3
	var arr3 = [3]int{4, 5, 6}
	fmt.Println(arr3)
	//方式4，不固定长度
	var arr4 = [...]int{1, 2, 3, 4}
	fmt.Println(arr4)
	//方式5，不固定长度，给指定下标赋值
	var arr5 = [...]int{1: 1, 0: 2, 2: 3, 3: 4}
	fmt.Println(arr5)

	fmt.Println("=============================")

	//数组的遍历
	//方式1
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	//方式2
	for i, num := range arr1 {
		fmt.Println("数组下标", i, "数组元素", num)
	}

	//二维数组的定义及遍历
	var arr6 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	//方式1
	for i := 0; i < len(arr6); i++ {
		for j := 0; j < len(arr6[i]); j++ {
			fmt.Print(arr6[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("==============")
	//方式2
	for _, arr := range arr6 {
		for _, num := range arr {
			fmt.Print(num, "\t")
		}
		fmt.Println()
	}

	fmt.Println("值传递 arra[0] 修改前", arr1[0])
	updateArrV(arr1)
	fmt.Println("值传递 arra[0] 修改后", arr1[0])

	fmt.Println("地址传递 arra[0] 修改前", arr1[0])
	updateArrE(&arr1)
	fmt.Println("地址传递 arra[0] 修改后", arr1[0])
}

//数组默认使用的是值传递
func updateArrV(arr [3]int) {
	arr[0] = 10
}

//使用指针的地址传递
func updateArrE(arr *[3]int) {
	(*arr)[0] = 10
}
