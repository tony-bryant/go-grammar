package chapter5

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func Main55() {
	//函数间不可比较 因此不能作为map的key  可以转换成string放入map
	//声明接收参数为int 返回int
	var f func(int) int
	//赋值给函数  函数作为一种类型
	f = square
	//打印函数结果
	fmt.Println(f(3)) // "9"

	f = negative
	fmt.Println(f(3))     // "-3"
	fmt.Printf("%T\n", f) // "func(int) int"

	//根据参数列表和返回列表判断函数类型
	//f = product

	//传入一个add1函数
	//模板方法
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"

	//将startElement和endElement传入forEachNode中
	forEachNode(nil, startElement, endElement)
}

func square(n int) int {
	return n * n
}
func negative(n int) int {
	return -n
}
func product(m, n int) int {
	return m * n
}

func add1(r rune) rune {
	return r + 1
}

//定义pre和post处理函数
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
var depth int
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
