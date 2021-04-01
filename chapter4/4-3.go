package chapter4

import (
	"fmt"
	"sort"
)

func Main43() {
	//1 使用字面值创建map
	ages1 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages1)
	//2 使用make函数创建map
	ages2 := make(map[string]int)
	ages2["alice"] = 31
	ages2["charlie"] = 34
	fmt.Println(ages2)
	//使用delete删除元素
	delete(ages2, "alice")
	//删除不存在的元素不会报错
	delete(ages2, "test")
	fmt.Println(ages2)
	//默认返回0值 bob=0+1
	ages2["bob"] += 1
	fmt.Println(ages2)
	//禁止对map元素进行取址操作 扩容会导致地址发生变化

	//使用range遍历map
	//遍历顺序是随机的，每一次遍历都不一定相同
	for name, age := range ages2 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//显式的堆key的slice进行排序，保证每次输出结果一致
	//var names []string
	//指定slice大小
	names := make([]string, 0, len(ages2))
	for name := range ages2 {
		names = append(names, name)
	}
	//排序names
	sort.Strings(names)
	for _, name := range names {
		//o(1)的查询
		fmt.Printf("%s\t%d\n", name, ages2[name])
	}

	//可以对nil的map进行增删改查等操作
	//但是对nil的map插入元素会导致panic
	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"
	//存入map前需要创建map
	age, ok := ages2["bob1"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println("不存在该人")
	}
	str := fmt.Sprintf("%q", ages2)
	fmt.Println(str)
	//go语言中不存在set，可以利用map中key不重复实现set
}

//不能直接使用==比较两个map
//省略的声明方式  x和y都是map类型的
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

//创建一个map
var m = make(map[string]int)

//将slice转换成字符串
func k(list []string) string { return fmt.Sprintf("%q", list) }

//向map元素增加计数
func Add(list []string) { m[k(list)]++ }

//返回计数的次数
func Count(list []string) int { return m[k(list)] }
//创建一个key为string value为map的map
var graph = make(map[string]map[string]bool)
