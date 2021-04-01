package chapter4

import (
	"fmt"
)

//声明一个结构体 结构体大写 属性开头大写包外可访问
//可以包含导出和未导出的成员
type Employee struct {
	ID int
	//可以合并相同类型的属性
	Name, Address string
	//Name      string
	//Address   string
	Position  string
	Salary    int
	ManagerID int
	Score     int //缺考还是0分？
}

//空结构体
//可以用在set的map声明中
type empty struct{}

//nil
var set map[string]struct{}

//结构体不可以包含本身 但是可以包含该类型的指针类型
type tree struct {
	value       int
	left, right *tree //指向tree的指针 *(&x)就是当前的值
}

//实体类型
//所有属性默认初始化 可以直接使用
var dilbert Employee

func Main44() {
	//对实体赋值
	dilbert.ID = 1
	fmt.Println(dilbert)

	//取址操作
	position := &dilbert.Position
	//通过指针操作
	*position = "Senior " + *position
	fmt.Println(dilbert)

	//语句相等
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Println(dilbert)

	tmp1 := Employee1()
	fmt.Println(tmp1)
	tmp1.Address = "test"
	fmt.Println(tmp1)
	tmp2 := Employee2()
	fmt.Println(*tmp2)
	(*tmp2).Address = "test"
	fmt.Println(*tmp2)

	//结构体字面量
	//写法一：顺序为每个字面量指定一个值
	employee1 := Employee{1, "gang", "ChenDu", "SC", 100, 1, 1}
	fmt.Println(employee1)
	//写法二：成员名和对应值进行初始化 顺序并不重要
	employee2 := Employee{Name: "huang", ID: 2}
	fmt.Println(employee2)

	//结构体可以作为函数的参数和返回值
	//可以通过结构体指针提高效率
	//只有使用指针才可以在函数体内部修改实体 否则修改的是值拷贝
	fmt.Println(*Employee3(&employee1))

	//结构体比较
	//结构体可以直接使用==进行比较
	//结构体可以作为map的key

	//结构体嵌入匿名成员 减少链式调用的复杂度
	//只声明数据类型而不指定成员名字
	//不能包含两个类型相同的匿名成员
	//简短的语法糖
	type Point struct {
		X, Y int
	}
	type Circle1 struct {
		Center Point//正常声明
		Radius int
	}
	type Circle2 struct {
		Point//匿名成员 减少链式调用 可以在引用时省略
		Radius int
	}
	//%#v将打印每个成员的名字
}

func Employee1() Employee {
	var tmp Employee
	tmp.ID = 1
	return tmp
}
func Employee2() *Employee {
	var tmp Employee
	tmp.ID = 2
	return &tmp
}
//接收指针，返回指针
func Employee3(tmp *Employee) *Employee {
	tmp.ID = 2
	return tmp
}
