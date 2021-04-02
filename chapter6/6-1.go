package chapter6

import (
	"fmt"
	"math"
)

//创建新类型 共用底层数据结构
type Path []Point

func Main61() {
	p:=Point{1,1}
	q:=Point{2,2}
	fmt.Println(Distance(p, q))
	//在p实体中调用方法Distance
	//其他类型实体不可调用
	fmt.Println(p.Distance(q))
	//p.Distance的表达式叫做选择器是 可以选择实体的类型和方法
	fmt.Println(p.X)

	//定义一个三角形
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	//计算三角形周长
	fmt.Println(perim.Distance()) // "12"
}
//定义结构体
//type Point struct{
//	X, Y float64
//}
//定义普通函数
func Distance(p, q Point) float64 {
	fmt.Println("通用Distance方法")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
//定义方法
//p 方法接收器  类似其他语言的this和self
//java中默认生成一个该类型的方法
func (p Point) Distance(q Point) float64 {
	fmt.Println("在Point类型中的Distance方法")
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
//指定接收path类型 path类型是关于Point Slice的命名类型
func (path Path) Distance() float64 {
	fmt.Println("在Path类型中的Distance方法")
	sum := 0.0
	//index value 灵活
	for i := range path {
		//忽略0开始
		if i > 0 {
			//此时调用在Point类型中的Distance方法
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
