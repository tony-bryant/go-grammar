package chapter7

import "io"

func Main75() {
	var w io.Writer
	w.Write([]byte("hello")) // panic: nil pointer dereference  编译通过
}
