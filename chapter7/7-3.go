package chapter7

import (
	"bytes"
	"fmt"
)

func Main73() {
	//空接口表示任意类型
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)

	fmt.Println(any)
}
