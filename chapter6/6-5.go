package chapter6

import (
	"bytes"
	"fmt"
)

func Main65() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}"
	//会隐式地在x前插入&操作符
	//隐式转换成指针调用
	fmt.Println(x.String()) // "{1 9 42 144}"
	//打印地址 &符号不能省略
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
}

//bit数组
type IntSet struct {
	//声明uint64切片  至多64个 懒加载
	words []uint64
}
//判断是否存在
//指针类型方法可以修改数组
//该位为1则存在
//指针用来配合Add和UnionWith
func (s *IntSet) Has(x int) bool {
	//数组下标 具体位置
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
//增加元素
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
//联合
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
//string方法 属于interface{}
//指针用来配合Add和UnionWith
//不建议绑定到指针上
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
