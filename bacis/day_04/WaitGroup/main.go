package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"http://www.huawei.com",
	"http://www.jd.com",
}

func main() {

	for _, url := range urls {
		// 每一个Url 启动一个goroutine,同时给wg+1
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Request.Host, resp.Status)
			}
		}(url)
	}
	//	 等待所有请求结束
	wg.Wait()
}
