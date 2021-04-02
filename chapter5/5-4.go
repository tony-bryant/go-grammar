package chapter5

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//func Main54() {
//
//}

//错误处理策略
//1.错误传播  子程序失败会变成该函数的失败
//1.1构建新的错误信息
func returnError1() ([]string, error) {
	//失败原因不止一种
	_, err := http.Get("https://www.baidu.com/")
	//是否失败
	if err != nil {
		return nil, err
	}
	return nil, err
}

//1.2重新构建失败信息
func returnError2() ([]string, error) {
	url := "https://www.baidu.com/"
	resp, err := http.Get(url)
	_, err = html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		//格式化错误信息
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return nil, err
}

//2.重新尝试失败操作
func WaitForServer(url string) error {
	//定义超时间隔
	const timeout = 1 * time.Minute
	//计算超时时间
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		//操作成功
		if err == nil {
			return nil // success
		}
		//打印失败日志
		log.Printf("server not responding (%s);retrying…", err)
		//休眠后 重试
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	//return Error
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

//3.输出错误已并结束程序 在main中使用
func Main54() {
	url := "https://www.baidu.com/"
	if err := WaitForServer(url); err != nil {
		//打印并退出
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)

		//简写形式
		log.Fatalf("Site is down: %v\n", err)
	}
}

//4.输出错误信息
//if err := Ping(); err != nil {
//log.Printf("ping failed: %v; networking disabled",err)
//fmt.Fprintf(os.Stderr, "ping failed: %v; networking disabled\n", err)
//}

//5.忽略错误
//dir, err := ioutil.TempDir("", "scratch")
//if err != nil {
//仅仅打印错误
//return fmt.Errorf("failed to create temp dir: %v",err)
//}
//// ...use temp dir…
//os.RemoveAll(dir) // ignore errors; $TMPDIR is cleaned periodically

//任何由文件结束引起的错误
var EOF = errors.New("EOF")

func test() ([]rune, error) {
	var result = []rune{}
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		//判断错误类型
		if err == io.EOF {
			break // finished reading
		}
		if err != nil {
			return nil, fmt.Errorf("read failed:%v", err)
		}
		result = append(result, r)
	}
	return result, nil
}
