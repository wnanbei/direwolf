package direwolf

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func newTestResponseServer() *httptest.Server {
	respString := `<html lang="zh-CN">
	<head>
	<title>Direwolf</title>
	</head>
	<body>
	<li><a href="/convenient/">is the most convenient</a></li>
	<li><a href="/easy/">and easy to use http client with Golang</a></li>
	<li><a href="/author/">南北</a></li>
	<li><a href="/time/">2019-06-21</a></li>
	</body>
	</html>`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check method is GET before going to check other features
		if r.Method != "GET" {
			log.Fatalf("Expected method %q; got %q", "GET", r.Method)
		}
		if r.URL.Path == "/" {
			w.Write([]byte(respString))
		}
		if r.URL.Path == "/GBK" {
			content, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(respString))
			w.Write(content)
		}
		if r.URL.Path == "/GB18030" {
			content, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(respString))
			w.Write(content)
		}
		if r.URL.Path == "/latin1" {
			content, _ := charmap.ISO8859_1.NewEncoder().Bytes([]byte(`<li><a href="/author/">...</a></li>`))
			w.Write(content)
		}
	}))
	return ts
}

func TestReExtract(t *testing.T) {
	ts := newTestResponseServer()
	defer ts.Close()

	resp, err := Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	result1 := resp.Re(`\d{4}-\d{2}-\d{2}`)
	if result1[0] != "2019-06-21" {
		t.Fatal("Response.Re() failed.")
	}
	t.Log("Response.Re() passed.")

	result2 := resp.ReSubmatch(`<a href.*?>(.*?)</a>`)
	if len(result2) != 4 {
		t.Fatal("Response.ReSubmatch() failed.")
	}
	if result2[3] != "2019-06-21" {
		t.Fatal("Response.ReSubmatch() failed.")
	}
	t.Log("Response.ReSubmatch() passed.")
}

func TestCssExtract(t *testing.T) {
	ts := newTestResponseServer()
	defer ts.Close()

	resp, err := Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	result1 := resp.CSS(`a`).First().Text()
	if result1 != "is the most convenient" {
		t.Fatal("Response.CSS().First().Text() failed.")
	}
	t.Log("Response.CSS().First().Text() passed.")

	result2 := resp.CSS(`body`).CSS(`li`).CSS(`a[href=\/time\/]`).First().Text()
	if result2 != "2019-06-21" {
		t.Fatal("Response.CSS() failed.")
	}
	t.Log("Response.CSS() passed.")

	result3 := resp.CSS(`a`).First().Attr("href")
	if result3 != "/convenient/" {
		t.Fatal("Response.CSS().First().Attr() failed.")
	}
	t.Log("Response.CSS().First().Attr() passed.")

	result4 := resp.CSS(`a`).First().AttrOr("noExists", "none")
	if result4 != "none" {
		t.Fatal("Response.CSS().First().AttrOr() failed.")
	}
	t.Log("Response.CSS().First().AttrOr() passed.")

	result5 := resp.CSS(`a`).At(2).Text()
	if result5 != "南北" {
		t.Fatal("Response.CSS().At() failed.")
	}
	t.Log("Response.CSS().At() passed.")
}

func TestResponseEncoding(t *testing.T) {
	ts := newTestResponseServer()
	defer ts.Close()

	resp3, err := Get(ts.URL + "/latin1")
	if err != nil {
		t.Fatal(err)
	}
	resp3.Encoding = "latin1"
	result3 := resp3.ReSubmatch(`<a href="/author/">(.*?)</a>`)
	// t.Log(resp3.Text())
	if result3[0] != "..." {
		t.Fatal("Response latin1 failed.")
	}
	t.Log("Response latin1 passed.")

	resp, err := Get(ts.URL + "/GBK")
	if err != nil {
		t.Fatal(err)
	}
	resp.Encoding = "GBK"
	result1 := resp.ReSubmatch(`<a href="/author/">(.*?)</a>`)
	if result1[0] != "南北" {
		t.Fatal("Response GBK failed.")
	}
	t.Log("Response GBK passed.")

	resp2, err := Get(ts.URL + "/GB18030")
	if err != nil {
		t.Fatal(err)
	}
	resp2.Encoding = "GB18030"
	result2 := resp2.ReSubmatch(`<a href="/author/">(.*?)</a>`)
	if result2[0] != "南北" {
		t.Fatal("Response GB18030 failed.")
	}
	t.Log("Response GB18030 passed.")
}
