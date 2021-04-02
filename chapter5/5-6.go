package chapter5

import (
	"fmt"
	"sort"
)

func Main56() {
	//记录了每一次函数的状态
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

	//breadthFirst(crawl, os.Args[1:])
}

//返回了一个匿名函数
//匿名函数输入为空 返回一个int
var x int

func squares() func() int {
	var y int
	return func() int {
		x++
		y++
		fmt.Println(x, y)
		return x * x
	}
}
//key当前课程 value前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//深度优先遍历搜索整个图
func topoSort(m map[string][]string) []string {
	//声明结果切片
	var order []string
	//声明一个map string bool
	seen := make(map[string]bool)
	//声明函数visitAll
	var visitAll func(items []string)
	//使用匿名函数
	//一直向下寻找
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				//递归该函数
				visitAll(m[item])
				//使用声明的order slice
				//防止扩容改变地址
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	//排序key
	sort.Strings(keys)
	//调用visitAll函数
	visitAll(keys)
	return order
}
//广度优先遍历
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				//生成一个待访问切片worklist 然后递归
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
//var links=[]string{}
//func crawl(url string) []string {
//	fmt.Println(url)
//	list, err := links.Extract(url)
//	if err != nil {
//		log.Print(err)
//	}
//	return list
//}


