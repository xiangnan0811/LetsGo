package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, er := httputil.DumpResponse(resp, true)
	if er != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}

// http
/*
1. 使用http客户端发送请求
2. 使用http.Client控制请求头部等
3. 使用httputil简化工作
*/

// http服务器的性能分析
/*
1. import _ "net/http/pprof"
2. 访问/debug/pprof/
3. 使用go tool pprof分析性能
*/
