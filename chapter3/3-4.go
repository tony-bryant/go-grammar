package chapter3

import "fmt"

func Main34() {
	c := 'a'
	//&&优先级高于||
	if 'a' <= c && c <= 'z' ||
		'A' <= c && c <= 'Z' ||
		'0' <= c && c <= '9' {
		fmt.Print("ok!")
	}
}
