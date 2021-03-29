package chaper2

import (
	"fmt"
)

func Main23() {
	PointerTest()
}

func variable() {
	var s string
	fmt.Println(s) // ""

	var i, j, k int                 // int, int, int
	var b, f, z = true, 2.3, "four" // bool, float64, string

	fmt.Printf(s)
	fmt.Printf("%d", i)
	fmt.Printf("%d", j)
	fmt.Printf("%d", k)
	fmt.Printf("%v", b)
	fmt.Printf("%v", f)
	fmt.Printf(z)
}

//public 供包外使用
//func ShortVariable()  {
//	f, err := os.Open(infile)
//	f, err := os.Create(outfile) // compile error: no new variables
//}

func PointerTest() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	//返回一个变量地址
	fmt.Println(*f())
	*f()=3
	fmt.Println(*f())
}

//返回一个指针类型  指向变量地址
func f() *int {
	v := 1
	return &v
}

