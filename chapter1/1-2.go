package chaper1

//gofmt会根据字母顺序进行排序
//方式1 列表形式 常用
import (
	"fmt"
	"os"
	"strings"
)

//方式2 独立写 不常用
//import "fmt"
//import "os"

//获取输入数据
//1.从命令行获取参数
//os包提供与操作系统交互的函数与变量
//os.Args程序的命令参数，由string的动态数组组成
//os.Args[0]是命令本身的名字
//os.Args[1:len(os.Args)]具体参数
func Main12() {
	//声明变量名 变量类型
	//隐式赋予其类型零值
	//数值形类型为0
	//字符串类型为""
	var s, sep string
	var num int
	//循环获取参数
	//i := 1 短变量声明 可以直接根据初始值类型省略类型声明
	//i++ i += 1 i = i + 1等价
	//但是++i --i是非法的
	for i := 1; i < len(os.Args); i++ {
		//+字符串连接符
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(num)
	Loop1()
	loop2()
	loop3()
	loop4()
	declareVariables()
}

func Loop1() {
	//for循环操作1
	//初始值  条件  操作
	//初始值 可以省略
	//条件 判断为true继续循环，判断为false结束循环 可以省略
	//操作 可以省略
	//for initialization; condition; post {
	//	// zero or more statements
	//}
	//for {
	//	// zero or more statements
	//}
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		//+字符串连接符
		s += sep + os.Args[i]
		sep = " "
	}
}

func loop2() {
	//for循环操作2
	//使用range
	//短变量声明
	s, sep := "", ""
	//从切片第一个开始遍历
	//_表示索引，没有用到 go语言不允许使用无用局部变量
	//arg列表值
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func loop3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func loop4() {
	//合并slice的string
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func declareVariables() {
	//短变量声明 不需要声明变量类型
	s := ""
	//默认零值初始化
	var _ string
	//较少使用
	var _ = ""
	//显式指明变量类型和名称
	var _ string = ""
	fmt.Println(s)
}
