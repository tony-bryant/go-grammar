package chapter6

import "fmt"

//实体. = 接收器
func Main62() {
	//使用方式1
	//p是指向Pointer(内存地址)的指针
	p1 := &Point{1, 2}
	p1.ScaleBy(2.0)
	fmt.Println(*p1)

	//使用方式2
	p2 := Point{1, 2}
	pptr := &p2
	//可以写成int 小转大不丢失精度
	pptr.ScaleBy(2)
	(*pptr).ScaleBy(2)
	fmt.Println(p2)

	//使用方式3
	//隐式使用&p3调用方法
	//只适用于变量 常量则会失败
	//Point{1, 2}.ScaleBy(2)
	p3 := Point{1, 2}
	p3.ScaleBy(2)
	fmt.Println(p3)

	//nil是合法的接收器类型
	//生成一个类型为Point的实体 占用内存空间
	//实体无值
	var p4 Point
	p4.ScaleBy(2)
}

//获取p的指针 可以直接修改实体的值
//改变实际的值
func (p *Point) ScaleBy(factor float64) {
	if p == nil {
		return
	}
	p.X *= factor
	p.Y *= factor
	//fmt.Println(*p)
}
