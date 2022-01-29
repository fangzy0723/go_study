package main

import (
	"fmt"
)

//map（映射） 类似于java中的hashMap 键值对关联
//基本语法  var mapName map[keyType]valueType
//map在使用前一定要mark(map[keyType]valueType)
//map的key-value是无须的，key不能重复 value可以重复
//key不能是 slice、map、function

func main() {
	//声明一个map
	var map1 map[string]string
	//对map进行初始化
	map1 = make(map[string]string)
	//赋值
	map1["name"] = "张三"
	fmt.Println(map1)

	//修改key的值,key存在是修改不存在是添加
	map1["name"] = "李四"
	fmt.Println(map1)

	//根据key查询value  查询返回两个值 1、key对应的value值  2、key是否存在   true存在  false不存在
	value, flag := map1["name"]
	fmt.Printf("value : %v 是否存在：%v \n", value, flag)
	//删除map中的元素
	delete(map1, "name")
	fmt.Println(map1)

	//方式2 声明并初始化
	map2 := make(map[string]string)
	map2["name"] = "haha"
	fmt.Println(map2)

	//方式3  初始化并赋值
	map3 := map[string]string{
		"name":  "zhangsan",
		"title": "haha", //, 不能少
	}
	fmt.Println(map3)

	//遍历 只能是for range
	for key, value := range map3 {
		fmt.Printf("key:%v value:%v \n", key, value)
	}

	//key是string  value是map类型  map的key是string，value是string类型
	map4 := make(map[string]map[string]string)
	map4["banji1"] = make(map[string]string)
	map4["banji1"]["name"] = "张三"
	map4["banji1"]["title"] = "haha"

	map4["banji2"] = make(map[string]string)
	map4["banji2"]["name"] = "张三"
	map4["banji2"]["title"] = "haha"

	for key, value := range map4 {
		fmt.Println("班级：", key)
		for k1, v1 := range value {
			fmt.Printf("key:%v value:%v \n", k1, v1)
		}
	}
}
