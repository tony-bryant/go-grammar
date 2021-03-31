package chapter3

import (
	"fmt"
	"math"
)

func Main32() {
	//超出有效位数23位
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %g\n", x, math.Exp(float64(x)))
		fmt.Printf("x = %d e^x = %f\n", x, math.Exp(float64(x)))
		fmt.Printf("x = %d e^x = %e\n", x, math.Exp(float64(x)))
	}

	//nan与任何数都不相等
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
}

func test1() (value float64, ok bool) {
	return 0, true
}

func test2() (float64, bool) {
	return 0, true
}
