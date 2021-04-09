package chapter7

import (
	"flag"
	"fmt"
	"github/gang/go-grammar/chapter2"
	"time"
)

//使用flag.Value定义新的符号
//更友好的定义输入形式
//定义命令行启动
var period = flag.Duration("period", 1*time.Second, "sleep period")

func Main74() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	//休眠
	time.Sleep(*period)
	//休眠结束
	fmt.Println()

	var test MyValue = MyName{name: ""}
	fmt.Println(test.String())
}

type MyValue interface {
	String() string
	Set(string) error
}

type MyName struct {
	name string
}

func (m MyName) String() string {
	return "huang gang"
}

func (m MyName) Set(name string) error {
	m.name = name
	return nil
}

type celsiusFlag struct{ chapter2.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = chapter2.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = chapter2.FToC(chapter2.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
