package chapter3

import "fmt"

func Main31() {
	//有符号数会使用符号位填充
	//无符号数左移和右移
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	//打印至少8个独立的bit位
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Println("---------")   // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	//len()返回有符号整数
	//无符号数只适合位运算和特殊场景中使用
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}

	//不同类型int
	var apples int32 = 1
	var oranges int16 = 2
	//编译错误
	//var compote int = apples + oranges // compile error
	//类型转换
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)

	//类型转换造成的精度丢失
	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	//字符字面
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	//[1]再次使用第一个操作数
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
