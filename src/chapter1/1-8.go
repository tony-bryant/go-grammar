package main

import "fmt"

//不需要显式声明break
//执行完成一个case后自动退出
//switch中可以使用
//1.简短的变量声明
//2.一个自增表达式
//3.赋值语句
//4.函数调用
func switchFun1(num int) {
	switch num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}
}

//关键字fallthrough会破坏执行完成后自动退出
func switchFun2(num int) {
	switch num {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("default")
	}
}

func loop() {
	stringSlice := []string{"test1", "test2", "test3"}
loop1:
	for index, string := range stringSlice {
		if index == 10 {
			continue
		} else if index == 11 {
			break
		} else {
		loop2:
			for _, c := range string {
				if c == '!' {
					continue loop1
				} else if c == '*' {
					break loop2
				} else {
					fmt.Printf("%c", c)
				}
			}
		}
	}
}

//声明结构体 type xxx struct
type Point struct {
	X, Y int
}

//指针 直接存储变量内存地址的数据类型
//可见的内存地址
//&操作返回变量的内存地址  实际地址
//*操作可以获取指针的变量内容  引用

//方法与接口
//接口是抽象类型
//实现类不需要显式的声明接口类型，方法一致则会自动对应


//单行注释
/*
多行注释
1
2
3
*/
