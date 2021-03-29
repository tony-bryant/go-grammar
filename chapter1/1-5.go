package chaper1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//通过net包网络获取数据
//无返回值则空着，不需要void
func Main15() {
	//声明一个地址切片
	var URLs = []string{"http://www.baidu.com", "http://www.qq.com", "http://www.sina.com"}
	start := time.Now()
	for _, url := range URLs {
		//创建HTTP的GET函数
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			//退出程序状态码为1 默认退出系统的状态码为2 正常状态状态码为0
			os.Exit(1)
		}
		//更新err的值
		//response.Body可读的响应流
		_, err = ioutil.ReadAll(response.Body)
		//关闭资源流，防止泄露
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		//fmt.Printf("%s", result)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs", secs)
}
