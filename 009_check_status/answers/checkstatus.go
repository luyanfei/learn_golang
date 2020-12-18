package main

import (
	"fmt"
	"net/http"
)

/*
checkStatus函数内部的goroutine在调用http.Get()方法时，在碰到异常时，选择了在控制台打印异常。
这是无奈的选择，这会导致调用方无从获得错误信息。
请改写代码，满足如下要求：
1. 能够在checkStatus外部处理错误，并正确打印出http响应的状态码。
2. 若连续出现三次错误，则打印出"Too many errors, breaking!"，并不再发送http请求。
(Concurrency in Go, ch4, Error Handling)
*/
type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{err, resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)
	urls := []string{"https://www.baidu.com", "https://badhost", "a", "b", "c", "d"}
	errorCount := 0
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v\n", result.Error)
			errorCount++
			if errorCount == 3 {
				fmt.Println("Too many errors, breaking!")
				break
			}
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}
