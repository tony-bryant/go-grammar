package chapter5

import (
	"fmt"
	"os"
	"runtime"
)

func Main59() {
	s := "Spades"
	switch s {
	case "Spades": // ...
	case "Hearts": // ...
	case "Diamonds": // ...
	case "Clubs": // ...
	//直接出发panic
	default:
		panic(fmt.Sprintf("invalid suit %q", s))
	}
	//输出堆栈信息
	defer printStack()
	f(3)
}

//x=0时发生panic异常
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
