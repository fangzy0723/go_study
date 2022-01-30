package main

//结构体的声明及实例的创建
//go 中没有类的概念   结构体类似于java中的类

import "fmt"

// Person 定义一个结构体 名为Person 属性为Name
type Person struct {
	Name string
	Age  int
}

// Student 给Person 定义一个别名
type Student Person

//方法的定义 格式： func (变量名 方法绑定的结构体类型) 方法名(参数列表) (返回值列表) {}
//方法默认是值传递，需要地址传递可使用指针（方法绑定的结构体使用指针类型）
//方法名首字母大写 可以其他包中使用，小写只能在本包中使用
//方法和结构体绑定后只能用这个结构题的实例调用
func (p Person) test1(str string) string {
	//默认值传递，在这赋值后原对象的属性值不变
	p.Name = "你哈"
	return p.Name + str
}
func (p *Person) test2(str string) string {
	//类型是指针 地址传递，会改变原对象的属性值
	p.Name = "你hao"
	return p.Name + str
}

//类型实现的String方法，fmt.Println(p) 会自动调用这个类型的String方法，类似于java中重写了toString方法
func (p Person) String() string {
	return fmt.Sprintf("Name:%v ,Age: %v", p.Name, p.Age)
}

//函数和方法的区别
//1、定义的时候 函数不需要绑定结构体类型，方法需要
//2、调用的时候 函数可以追调用，方法需要用对象.的方式： 结构体对象.方法名(参数列表)
//3、函数调用参数到类型和形参类型一致，方法调用参数是指针类型可以传值，值类型可以传指针

func main() {
	//结构体实例的创建方式
	//方式1
	var p1 Person = Person{
		Name: "张三",
		Age:  20,
	}
	fmt.Println(p1)

	var p3 Person = Person{"张哈", 20} //值的位置和结构体中属性的位置一一对应
	fmt.Println(p3)

	var p2 Person
	p2.Name = "Lisa"
	fmt.Println(p2.Name)

	//方式2
	//p4是指针类型  指向这个实例的内存地址
	var p4 *Person = new(Person)
	//给属性直接赋值（如下是指针的简化赋值方式）
	p4.Name = "丽萨"
	p4.Age = 25
	fmt.Println(p4)

	//方式3
	var p5 *Person = &Person{"haha", 10}
	fmt.Println(*p5)
	fmt.Println(p5.Name)

	//结构体之间转换，属性名名称和类型、顺序一直可以转换
	var p6 Person = Person{"张哈", 20}
	var s1 Student = Student{"换手机", 20}
	p6 = Person(s1)

	fmt.Println(s1.Name)
	fmt.Println(p6.Name)

	fmt.Println("============")
	fmt.Println(p6.test1("haha"))
	fmt.Println(p6.Name)
	fmt.Println((&p6).test2("hehe"))
	fmt.Println(p6.test2("222")) //方法调用参数是指针类型可以传值，值类型可以传指针 会自动进行隐式转换
	fmt.Println(p6.Name)
	fmt.Println("--------------")

	//Person绑定了String方法，打印对象会自动调用该对象的String方法
	fmt.Println(p6)

}
