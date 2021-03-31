package chapter3

import (
	"fmt"
	"math"
)

const (
	a = 1
	b
	c = 2
	d
)

//iota常量生成器
//类似枚举类型
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

//go语言没有计算幂的运算符
const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776             (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424    (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

const pi = 3.14159

func Main36() {
	fmt.Println(a, b, c, d) // "1 1 2 2"
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(FlagUp)        //00001
	fmt.Println(FlagBroadcast) //00010
	fmt.Println(FlagLoopback)  //00100

	//进行响应类型转换
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
}
