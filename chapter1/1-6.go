package chaper1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//并发获取URL内容  goroutine+channel
//主goroutine
//当goroutine在channel上做send或者receive操作时，这个goroutine会阻塞在调用处
func Main16() {
	//声明一个地址切片
	var URLs = []string{"http://www.baidu.com", "http://www.qq.com", "http://www.sina.com"}
	//创建不限长的通道
	ch := make(chan string)
	for _, url := range URLs {
		//创建并开启一个goroutine并发执行方式
		go fetch(url, ch)
	}
	//只需要获取URLs个数
	for range URLs {
		//获取通道内容时会阻塞，直到通道写完成
		//依次打印通道内容
		//避免异步执行未完成后提前退出 类似countDownLatch
		fmt.Println(<-ch)
	}
}

//所有类型小写
func fetch(url string, ch chan<- string) {
	//获取开始时间
	start := time.Now()
	//异步执行GET方法
	response, err := http.Get(url)
	if err != nil {
		//将错误写入通道
		ch <- fmt.Sprint(err)
		return
	}
	_, err = ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		//字符串格式化指令
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}
	//执行时间
	secs := time.Since(start).Seconds()
	//将结果写入通道中
	//会阻塞到别的goroutine调用通道内容后，继续向下执行
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}
