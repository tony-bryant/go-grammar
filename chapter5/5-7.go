package chapter5

import (
	"fmt"
	"os"
)

func Main57() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"

	//将切片类型传入sum函数 需要加上...
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

//参数列表是一个数组
func sum(vals...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

//接收args可变参数
//interface{}可以接收任意类型
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
