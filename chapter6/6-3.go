package chapter6

import (
	"fmt"
	"image/color"
	"sync"
)

func Main63() {
	var cp ColoredPoint
	cp.X = 1
	//显式链式调用
	fmt.Println(cp.Point.X) // "1"
	cp.Point.Y = 2
	//匿名成员调用
	fmt.Println(cp.Y) // "2"

	//声明两个color实体
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	//声明ColoredPoint实体
	//使用var和短声明都可以
	p:= ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	//使用ColoredPoint当做接收器进行方法调用
	//类似Point是基类 ColoredPoint是子类
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
}

//定义普通结构体
//包外可使用
//结构体与结构体方法声明可以不在一个go文件中
type Point struct{ X, Y float64 }

//嵌套Point的结构体
type ColoredPoint struct {
	//声明匿名成员 不需要长的链式调用
	Point
	//也可以声明为指针
	//*Point
	Color color.RGBA
}

//可以访问ColoredPoint其他字段
//优先使用这个方法
//同一级存在同名方法则报错
func (p ColoredPoint) Distance(q Point) float64 {
	fmt.Println("在ColoredPoint类型中Distance方法")
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	fmt.Println("在ColoredPoint类型中ScaleBy方法")
	p.Point.ScaleBy(factor)
}

//新的结构体变量取别名
var cache = struct {
	//匿名成员
	sync.Mutex
	mapping map[string]string
}{
	//创建这个map
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	//匿名成员sync.Lock()
	cache.Lock()
	//使用mapping的get
	v := cache.mapping[key]
	//匿名成员sync.Unlock()
	cache.Unlock()
	return v
}
