package direwolf

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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

func TestParams(t *testing.T) {
	URL := "http://test.com"
	params := NewParams(
		"key1", "value2",
		"key2", "value2")
	req, err := NewRequest("GET", URL, params)
	if err != nil {
		t.Fatal(err)
	}
	if req.URL != "http://test.com?key1=value2&key2=value2" {
		t.Fatal("Test params failed.")
	}

	URL = "http://test.com?"
	params = NewParams(
		"key1", "value2",
		"key2", "value2")
	req, err = NewRequest("GET", URL, params)
	if err != nil {
		t.Fatal(err)
	}
	if req.URL != "http://test.com?key1=value2&key2=value2" {
		t.Fatal("Test params failed.")
	}

	URL = "http://test.com?xxx=yyy"
	params = NewParams(
		"key1", "value2",
		"key2", "value2")
	req, err = NewRequest("GET", URL, params)
	if err != nil {
		t.Fatal(err)
	}
	if req.URL != "http://test.com?xxx=yyy&key1=value2&key2=value2" {
		t.Fatal("Test params failed.")
	}
}

func GetJsonTestServer() *httptest.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/json", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		c.String(200, string(data))
	})
	ts := httptest.NewServer(router)
	return ts
}

func TestJsonBody(t *testing.T) {
	ts := GetJsonTestServer()

	type Student struct {
		Name  string
		Age   int
		Guake bool
	}
	jsonText := &Student{
		"Xiao Ming",
		16,
		true,
	}

	jsonBody := NewJsonBody(jsonText)
	resp, err := Get(ts.URL+"/json", jsonBody)
	if err != nil {
		t.Fatal("TestJsonBody Failed.")
	}

	if resp.JsonGet("Name").String() != "Xiao Ming" {
		t.Fatal("TestJsonBody Failed.")
	}

	jsonRespBody := &Student{}
	if err := resp.Json(jsonRespBody); err != nil {
		t.Fatal("TestJsonBody Failed.")
	}
	if jsonRespBody.Name != "Xiao Ming" {
		t.Fatal("TestJsonBody Failed.")
	}
}
