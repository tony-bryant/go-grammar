package chapter6

import (
	"fmt"
	"time"
)

func Main64() {
	//声明两个实例
	p := Point{1, 2}
	q := Point{4, 6}

	//指定使用方法
	distanceFromP := p.Distance
	//调用该方法 以后不需要指定实体类型
	fmt.Println(distanceFromP(q))

	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", sqrt(5)
	fmt.Println(p.Distance(origin))    // "2.23606797749979", sqrt(5)

	//绑定实体与方法名 后面无需再次指定
	scaleP := p.ScaleBy // method value
	scaleP(2)           // p becomes (2, 4)
	scaleP(3)           // then (6, 12)
	scaleP(10)          // then (60, 120)

	//生成指向Rocket实体的指针
	r := new(Rocket)
	//定时后执行
	//传入匿名函数
	time.AfterFunc(10*time.Second, func() { r.Launch() })
	//传入函数 函数类型 +()是执行函数
	time.AfterFunc(10*time.Second, r.Launch)
	distance := Point.Distance   // method expression
	//方法表达式
	//第一参数是实体类型
	//第二参数是函数参数
	//func (p Point) Distance(q Point) float64
	fmt.Println(distance(p, q))  // "5"
}

type Rocket struct{ /* ... */ }

func (r *Rocket) Launch() { /* ... */ }

//声明两个Point方法
func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

func (path Path) TranslateBy(offset Point, add bool) {
	//声明方法表达式
	var op func(p, q Point) Point
	if add {
		//方式值
		op = Point.Add
	} else {
		//方式值
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
