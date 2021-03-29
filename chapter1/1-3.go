package chaper1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Main13() {
	//dup1()
	print()
}

func dup1() {
	//创建一个key=string value=int的map
	//键类型是可以通过==比较的类型 不包括函数类型、 字典类型、切片
	counts := make(map[string]int)
	//通过bufio创建输入
	//使用段声明创建，不需要指定其类型
	input := bufio.NewScanner(os.Stdin)
	//通过map计数输入
	for input.Scan() {
		//首次使用value=0
		//再次使用value为上次值
		//=getDefaultValue(map.get(key),0)
		counts[input.Text()]++
	}
	//利用range短变量声明line key、n value
	//range也可以用于slice
	for line, n := range counts {
		//条件不需要加括号，但是主体需要加括号
		if n > 1 {
			fmt.Print("%d\t%s\n", n, line)
		}
	}
}

func print() {
	//\t          制表符
	//\n          换行符
	//%d          十进制整数
	fmt.Printf("10=%d\n", 10)
	//%x, %o, %b  十六进制，八进制，二进制整数
	fmt.Printf("10=%x\n", 10)
	fmt.Printf("10=%o\n", 10)
	fmt.Printf("10=%b\n", 10)
	//%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
	fmt.Printf("3.141592653589793=%f\n", 3.141592653589793)
	fmt.Printf("3.141592653589793=%g\n", 3.141592653589793)
	fmt.Printf("3.141592653589793=%e\n", 3.141592653589793)
	//%t          布尔：true或false
	fmt.Printf("true=%t\n", true)
	fmt.Printf("false=%t\n", false)
	//%c          字符（rune） (Unicode码点)
	fmt.Printf("a=%c\n", 'a')
	//%s          字符串
	fmt.Printf("黄罡=%s\n", "黄罡")
	//%q          带双引号的字符串"abc"或带单引号的字符'c'
	fmt.Printf("黄罡=%q\n", "黄罡")
	//%v          变量的自然形式（natural format）
	fmt.Printf("黄罡=%v\n", "黄罡")
	//%T          变量的类型
	fmt.Printf("string type=%T\n", "黄罡")
	//%%          字面上的百分号标志（无操作数）
	fmt.Printf("%%\n")
	//Println会自动添加一个换行符
}

func dup2() {
	//声明map
	counts := make(map[string]int)
	//切分slice，获取新的slice
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		//_=索引 arg=slice的值
		for _, arg := range files {
			//f=文件 err=打开文件是否存在错误
			f, err := os.Open(arg)
			//判断是否存在错误
			if err != nil {
				//存在错误则打印错误信息
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				//执行下一个循环
				continue
			}
			countLines(f, counts)
			//关闭文件
			f.Close()
		}
	}
	for line, n := range counts {
		//条件不需要加括号，但是主体需要加括号
		if n > 1 {
			fmt.Print("%d\t%s\n", n, line)
		}
	}
}

//在其声明前被调用
//函数与包级变量可以在任意位置声明
func countLines(f *os.File, counts map[string]int) {
	//counts是创建的map指针副本  修改的仍然是dup2中的map
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

//一般使用bufio中的高级函数如io/ioutil
//将数据全部读入内存中
func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//获取字节切片
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//字节slice转换为string后进行切分
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
