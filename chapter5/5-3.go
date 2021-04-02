package chapter5

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func Main53() {
	links, err := findLinks("https://www.nowcoder.com/")
	//直接打印多个结果
	log.Print(findLinks("https://www.nowcoder.com/"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}

//返回一个string切片和error错误
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	//获取是否失败
	if err != nil {
		return nil, err
	}
	//判断http状态
	if resp.StatusCode != http.StatusOK {
		//该部分不会被垃圾回收 需要手动释放资源
		resp.Body.Close()
		//输出详细的错误信息
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	//fmt.Println(doc)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	//包内访问
	//s := make([]string,0)
	return visit(nil, doc), nil
}

//bare return
//减少代码量 增加代码理解难度
func operation(a int, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("分母不为0")
		return
	}
	result = a / b
	return
}
