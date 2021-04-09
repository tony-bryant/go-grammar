package chapter7

import (
	"bytes"
	"fmt"
	"os"
)

func Main71() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	//使用Fprintf
	fmt.Fprintf(&c, "hello, %s", name)
	//使用Println
	fmt.Println(c) // "12", = len("hello, Dolly")
}

//返回int error
func Printf(format string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stdout, format, args...)
}
//返回string
func Sprintf(format string, args ...interface{}) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, format, args...)
	return buf.String()
}
//定义一个接口
//type MyWriter interface {
//	Write(p []byte) (n int, err error)
//}

//func Fprintf(w io.Writer, format string, args ...interface{}) (int, error)

// 类型重命名
type ByteCounter int

//类型方法 可以使用Write方法
//满足Writer接口
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//声明String()方法
type MyStringer interface {
	String() string
}
