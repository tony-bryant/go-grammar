package chapter5

import (
	"fmt"
	"math"
)

func Main51() {
	//调用时不用指定参数
	fmt.Println(hypot1(3, 4))
	fmt.Printf("%T\n", hypot1) //打印函数标识符
	fmt.Println(hypot2(3, 4))
	fmt.Printf("%T\n", hypot2) //hypot1和hypot2是相同的函数
	//fmt.Println(hypot1 == hypot2)
}

//参数由函数调用者提供
//返回值变量名与类型
//可以省略相同参数类型
func hypot1(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func hypot2(x, y float64) (result float64) {
	return math.Sqrt(x*x + y*y)
}

//不是go语言实现的
//func Sin(x float64) float //implemented in assembly language
