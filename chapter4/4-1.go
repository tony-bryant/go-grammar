package chapter4

import "fmt"

func Main41() {
	var a [4]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	//index value
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	var b [3]int = [3]int{1, 2, 3}
	var c [3]int = [3]int{1, 2}
	fmt.Println(b[2])
	fmt.Println(c[2])

	//数组长度不同，代表不同的数据类型
	fmt.Printf("%T\n", a) //[4]int
	fmt.Printf("%T\n", b) //[3]int

	//不直接指定个数
	//短声明
	d := [...]int{1, 2, 3}
	fmt.Println(d[2])

	//索引与队列别表的初始值
	type Currency int
	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
	fmt.Printf("%T\n", symbol)    //数组  [4]string

	//最后一个元素初始化成-1，其余赋值为0
	r := [...]int{99: -1}
	fmt.Println(r[0])
	fmt.Println(r[99])

	//当两个数组类型和元素都相等时
	//==才会判断为true
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1) // "true false false"
	//d1 := [3]int{1, 2}
	//fmt.Println(a1 == d1) // compile error: cannot compare [2]int == [3]int

	//将数组指针传入
	zeroInt(&b)
	fmt.Println(b)
}

//接收指针类型
func zeroInt(ptr *[3]int) {
	for i := range ptr {
		ptr[i] = 0
	}
}
