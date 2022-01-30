package structTest

//定义一个私有的结构体，只能在本包中使用，别的包中需要使用该结构体只需要使用提供的公有方法
type person struct {
	name string
	age  int
}

// NewPerson 定义一个提供person结构体的函数，返回一个改结构体的实例指针
func NewPerson(name string, age int) *person {
	return &person{name, age}
}

// SetName 提供set方法，给属性设置值
func (p *person) SetName(name string) {
	p.name = name
}
func (p *person) SetAge(age int) {
	p.age = age
}

// GetName 提供get方法，获取属性的值
func (p *person) GetName() string {
	return p.name
}
func (p person) GetAge() int {
	return p.age
}
