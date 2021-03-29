package chapter1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)
//将结果输出到不同的地方 都实现了io.Writer
//输出结果到io.Writer
//输出结果到os.Stdout
//输出结果到文件
//输出结果到ioUtil.Discard
//输出结果到http.ResponseWriter

//互斥锁 将不同请求的goroutine串行化执行
//否则产生竞态条件
var mu sync.Mutex
//计数器
var count int

//web微服务 返回当前用户正在访问的URL
//每一个请求都启动一个goroutine  服务器可以同时处理多个请求
func Main17() {
	//将handler函数作为类型传入
	//将请求转发到handler函数中处理
	//类似模板方法
	//地址代表默认地址
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	//使用匿名函数
	//http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
	//	lissajous(w)
	//})
	//监听端口
	//打印日志
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//创建函数
func handler(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	//标准输出流
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	//判断前先赋值，然后判断
	//写法1 更优雅
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//写法2
	err := r.ParseForm()
	if err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	//标准输出流
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
