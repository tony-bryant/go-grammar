package chapter3

import "fmt"

func Main33() {
	var x complex128 = complex(1, 2) // 1+2i
	var y complex128 = complex(3, 4) // 3+4i
	fmt.Println(x * y)               // "(-5+10i)"
	fmt.Println(real(x * y))         // "-5"
	fmt.Println(imag(x * y))         // "10"

	//简化x，y的描述
	x1 := 1 + 2i
	y1 := 3 + 4i
	fmt.Println(x1 * y1)       // "(-5+10i)"
	fmt.Println(real(x1 * y1)) // "-5"
	fmt.Println(imag(x1 * y1)) // "10"

	//复数比较
	//实部与虚部相同时才相等
	fmt.Println(x1 == y1)
	fmt.Println(x1 != y1)
}
