package chapter3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Main35() {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')

	//左闭右开产生新的字符串
	fmt.Println(s[0:5]) // "hello"
	fmt.Println(s[:5])  // "hello"
	fmt.Println(s[7:])  // "world"
	fmt.Println(s[:])   // "hello, world"

	//比较字符串
	s1 := s[0:5]
	fmt.Println(s == s1)
	fmt.Println(s > s1)

	//增加字符串
	//产生新的字符串
	s += s1
	fmt.Println(s)

	//禁止修改字符
	//s[0] = 'L'

	//原生字符串
	const GoUsage = `Go is a tool for managing Go source code.

Usage:
    go command [arguments]
...`
	fmt.Println(GoUsage)

	s = "Hello, 世界"
	//字节数
	fmt.Println(len(s)) // "13"
	//字符数
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	//打印每个字符
	for i := 0; i < len(s); {
		//解析String
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for index, tmp := range s {
		fmt.Println(index, "    ", string(tmp))
	}
	r := []rune(s)
	//每个十六进制数字前插入一个空格
	fmt.Printf("%x\n", r)

	//不改变原字符串
	//仅改变拷贝的数组
	b := []byte(s)
	s2 := string(b)
	fmt.Println(b)
	fmt.Println(s2)
	b[0] = 'c'
	fmt.Println(b)
	fmt.Println(s2)

	fmt.Println(strings.Count(s1, "h"))

	//字符串与数字转换
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	//x, err = strconv.Atoi("123")             // x is an int
	//y, err = strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
}

//前缀测试
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

//后缀测试
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

//子串测试
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		//遍历每个起点
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

//返回文件名
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		//逆序寻找
		if s[i] == '/' {
			//重新赋值/后的字符串
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		//逆序寻找
		if s[i] == '.' {
			//保留.前的文件名
			s = s[:i]
			break
		}
	}
	return s
}

func basename1(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//递归切分
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
