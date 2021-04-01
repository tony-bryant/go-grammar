package chapter4

import "fmt"

func Main42() {
	//声明数组
	months := [...]string{1: "一月", 2: "二月", 3: "三月",
		4: "四月", 5: "五月", 6: "六月", 7: "七月", 8: "八月",
		9: "九月", 10: "十月", 11: "十一月", 12: "十二月"}
	//0为空字符串，从1开始有值
	fmt.Println(months[0])
	fmt.Println(len(months))
	fmt.Printf("%T\n", months)
	//slice切片，数组是定长的
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Printf("%T\n", Q2) //[]string
	fmt.Println(summer)
	fmt.Printf("%T\n", summer) //[]string
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				//比较字符串是否相等
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	str1 := "你好"
	str2 := "你好"
	//字符串比较字符是否相等
	fmt.Println(str1 == str2)
	Q2[0] = "测试修改"
	fmt.Println(months)    //不改变数组
	fmt.Println(Q2)        //改变切片
	fmt.Printf("%T\n", Q2) //[]string

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	//修改了数组底层
	fmt.Println(a) // "[5 4 3 2 1 0]"

	var s []int // len(s) == 0, s == nil
	fmt.Println(s == nil)
	s = nil // len(s) == 0, s == nil
	fmt.Println(s == nil)
	s = []int(nil) // len(s) == 0, s == nil
	fmt.Println(s == nil)
	s = []int{} // len(s) == 0, s != nil
	fmt.Println(s == nil)
	fmt.Println(len(s) == 0)

	//初始化slice方法
	//1 顺序指定初始化
	//2 索引和元素值指定初始化
	//3 混合语法初始化
	//4 make函数创建
	s2 := make([]int, 5, 10)
	fmt.Println(s2)
	//增加元素
	s2 = append(s2, 11)
	fmt.Println(s2)

	var x, y []int
	for i := 0; i < 10; i++ {
		//将结果复制给slice
		//runes = append(runes, r)
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t len=%d\t%v\n", i, cap(y), len(y), y)
		x = y
	}
	//一次append多个值
	x = append(x, 4, 5, 6)
	fmt.Printf("%d cap=%d\t len=%d\t%v\n", 10, cap(x), len(x), x)

	strings := []string{"aaa", "bbb", "","","", "ccc", ""}
	fmt.Println(strings)
	strings1 := nonempty(strings)
	//改变了原数组
	fmt.Println(strings)
	fmt.Println(strings1)
}

//传递slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//负责处理[]int类型的slice
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	//是否有足够容量保存新元素
	if zlen <= cap(x) {
		// 生成新的切片
		z = x[:zlen]
	} else {
		zcap := zlen
		//扩容两倍
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		//复制元素 此时数组与原数组不相同了
		//两个slice共享一个数组
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

//复制slice 传入的是引用
func nonempty(strings []string) []string {
	//短声明0，用于最后切分使用
	i := 0
	for _, s := range strings {
		if s != "" {
			//改变了原数组
			strings[i] = s
			i++
		}
	}
	//返回slice
	return strings[:i]
}
