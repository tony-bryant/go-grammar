//package go语言中的包，每个源文件都需要声明package
//可独立执行的程序，而不是一个库
package main

//导入别的包
//导入必须的包，删除未导入的包（佛者编译失败）
import "fmt"

//函数 func
//变量 var
//常量 const
//类型 type

func main() {
	test(1)
}

//func关键字  main函数名  a int参数列表 bool返回值类型 函数体
//自动将换行符转化成分号
//特殊符号不会换行
//gofmt 标准化格式
//goimports 自动添加或删除import声明
//1 标识符 整数 浮点数 虚数 字符 字符串文字
//2 break continue fallthrough return
//3 ++ -- ) ] }
func test(a int) bool {
	fmt.Println(a)
	return true
}
