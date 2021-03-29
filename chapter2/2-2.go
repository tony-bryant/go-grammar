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
func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func Test() float64 {
	return 0.2
}
