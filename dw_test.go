package direwolf

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func myCheckRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 2 {
		fmt.Println(1)
		return errors.New("stopped after 10 redirects")
	}
	return nil
}

// func TestHTTP(t *testing.T) {
// 	options := cookiejar.Options{
// 		PublicSuffixList: publicsuffix.List,
// 	}
// 	jar, err := cookiejar.New(&options)
// 	// 生成client客户端
// 	client := &http.Client{
// 		Jar: jar,
// 	}
// 	// 生成Request对象
// 	req, err := http.NewRequest("Get", "http://httpbin.org/cookies/set/hello/world", nil)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// 添加Header
// 	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36")
// 	// 发起请求
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// 设定关闭响应体
// 	defer resp.Body.Close()
// 	// 读取响应体
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(string(body))
// }

// func TestGet(t *testing.T) {
// 	h := Headers{
// 		"aaa": {"bbb", "ccc", "bbb"},
// 	}
// 	c := Cookies{
// 		"xxxx": "yyyyy",
// 		"qqqq": "w",
// 	}
// 	Get("http://httpbin.org/get", h, c)
// }

// func TestHTTP(t *testing.T) {
// 	headers := Headers{
// 		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36"},
// 	}
// 	params := Params{
// 		"keyxxx": {"valuexxx"},
// 	}
// 	resp := Get("http://httpbin.org/get", headers, params)
// 	mystr := "a"
// 	result := resp.ReSubmatch(mystr)
// 	fmt.Println(result)
// 	if result == nil {
// 		print("same")
// 	}
// }

type a struct {
	h []string
}

func TestRe(t *testing.T) {
	x := a{}
	if x.h == nil {
		print("same")
	}
	print(x.h)
}
