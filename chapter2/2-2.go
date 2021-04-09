package chapter2

import (
	"fmt"
)

//包一级范围声明 包内可用
const boilingF = 212.0

func Main22() {
	//函数内部声明 只能在函数内部是用
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, FToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, FToC(boilingF))   // "212°F = 100°C"
}

//包内可以调度用
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func Test() float64 {
	return 0.2
}
