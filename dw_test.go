package direwolf

import (
	"fmt"
	"testing"
)

func TestHTTP(t *testing.T) {
	headers := Headers{
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.108 Safari/537.36"},
	}
	params := Params{
		"keyxxx": {"valuexxx"},
	}
	var proxy Proxy = "http://127.0.0.1:12333"
	resp := Get("http://httpbin.org/get", headers, params, proxy)
	result := resp.ReSubmatch(`origin": "(.*?)"`)
	fmt.Println(result)
}
