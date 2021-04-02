package chapter5

import (
	"fmt"
	"golang.org/x/net/html"
)

func Main510() {

}
//deferred函数帮助从panic中恢复
func Parse(input string) (err error) {
	defer func() {
		//使用恢复函数
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	// ...parser...
	return err
}
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	//defer + 匿名函数
	defer func() {
		//恢复panic
		switch p := recover(); p {
		case nil:       // 无panic不作处理 继续向下执行
		case bailout{}: // 期望中的panic 抛出错误
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // 原封不动抛出panic
		}
	}()
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				//抛出panic错误
				panic(bailout{}) // multiple titleelements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
