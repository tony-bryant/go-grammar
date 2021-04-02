package chapter5

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
)

//声明Node结构体
type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

//重命名类型
type NodeType int32
//声明常量
const (
	ErrorNode NodeType = iota //使用iota递增
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

//func Parse(r io.Reader) (*Node, error)

func Main52() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	outline(nil, doc)
}
//返回所有的地址
func visit(links []string, n *html.Node) []string {
	//判断是否为a标签
	if n.Type == html.ElementNode && n.Data == "a" {
		//遍历a标签属性
		for _, a := range n.Attr {
			//判断是否为href
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	//递归每个子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	//返回结果切片
	return links
}

//动态调整大小
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		//入栈
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//没有使用指针 使用的是栈的复制 没有改变源的栈
		outline(stack, c)
	}
}
