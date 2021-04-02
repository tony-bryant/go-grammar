package chapter5

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
	"time"
)

func Main58() {
	bigSlowOperation()
	_ = double(4)//8
}
//需要处理的错误增加 维护逻辑越来越困难
//解决类似.Close()操作
func title1(url string) error {
	//get请求url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//获取Content-Type
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		//终止函数并抛出错误
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html",url, ct)
	}
	//解析responseBody
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title"&&n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}
func title2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	//使用defer关闭
	//return完成后逆序向上执行
	defer resp.Body.Close()
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct,"text/html;") {
		return fmt.Errorf("%s has type %s, not text/html",url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url,err)
	}
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title"&&n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}
	forEachNode(doc, visitNode, nil)
	return nil
}
func bigSlowOperation() {
	//trace("bigSlowOperation")调用函数
	//trace("bigSlowOperation")()执行匿名函数
	//defer会在return后再执行
	//在sleep完后再执行
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	//记录进入时间
	log.Printf("enter %s", msg)
	//打印结束时间
	return func() {
		log.Printf("exit %s (%s)", msg,time.Since(start))
	}
}
//使用defer函数观察匿名函数的结果
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x,result) }()
	return x + x
}
