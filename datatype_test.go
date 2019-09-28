package direwolf

import (
	"testing"
)

func TestHeaders(t *testing.T) {
	headers := NewHeaders(
		"key1", "value1",
		"key2", "value2",
	)
	if headers.Get("key1") != "value1" {
		t.Fatal("Headers.Get() failed.")
	}
	if headers.Get("key3") != "" {
		t.Fatal("Headers.Get() failed.")
	}
}

//func TestParams(t *testing.T) {
//	//params := NewParams("xxx", "yyy")
//	//proxy := &Proxy{
//	//	HTTPS: "xxx",
//	//	HTTP: "xxx",
//	//}
//	//req := NewRequestSetting("Post", "http://www.baidu.com", params, proxy)
//	//fmt.Printf("%+v", req)
//	option := &SessionOptions{}
//	*option.MaxIdleConns = 10
//	fmt.Print(option)
//}

func TestStringSliceMap(t *testing.T) {
	params := NewParams(
		"key1", "value1",
		"key2", "value2",
	)
	if params.Get("key1") != "value1" {
		t.Fatal("params.Get() failed.")
	}
	if params.Get("key3") != "" {
		t.Fatal("params.Get() failed.")
	}

	if params.URLEncode() != "key1=value1&key2=value2" {
		t.Fatal("params.URLEncode() failed.")
	}

	params.Add("key1", "value3")
	if params.Get("key1", 1) != "value3" {
		t.Fatal("params.Add() failed.")
	}

	params.Set("key1", "value4")
	if params.Get("key1") != "value4" {
		t.Fatal("params.Set() failed.")
	}

	params.Del("key2")
	if params.Get("key2") != "" {
		t.Fatal("params.Del() failed.")
	}

	postForm := NewPostForm(
		"key1", "value1",
		"key2", "value2",
	)
	if postForm.Get("key1") != "value1" {
		t.Fatal("PostForm.Get() failed.")
	}
}
