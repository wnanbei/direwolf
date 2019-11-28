package direwolf

import (
	"github.com/gin-gonic/gin"
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
	<li><a href="/convenient/">is a convenient</a></li>
	<li><a href="/easy/">and easy to use http client with Golang</a></li>
	<li><a href="/author/">南北</a></li>
	<li><a href="/time/">2019-06-21</a></li>
	</body>
	</html>`

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(200, respString)
	})
	router.GET("/GBK", func(c *gin.Context) {
		content, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(respString))
		c.Data(200, "text/html", content)
	})
	router.GET("/GB18030", func(c *gin.Context) {
		content, _ := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(respString))
		c.Data(200, "text/html", content)
	})
	router.GET("/latin1", func(c *gin.Context) {
		content, _ := charmap.ISO8859_1.NewEncoder().Bytes([]byte(`<li><a href="/author/">...</a></li>`))
		c.Data(200, "text/html", content)
	})
	ts := httptest.NewServer(router)
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

	result2 := resp.ReSubmatch(`<a href.*?>(.*?)</a>`)
	if len(result2) != 4 {
		t.Fatal("Response.ReSubmatch() failed.")
	}
	if result2[3][0] != "2019-06-21" {
		t.Fatal("Response.ReSubmatch() failed.")
	}
}

func TestCssExtract(t *testing.T) {
	ts := newTestResponseServer()
	defer ts.Close()

	resp, err := Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	result1 := resp.CSS(`a`).First().Text()
	if result1 != "is a convenient" {
		t.Fatal("Response.CSS().First().Text() failed.")
	}

	result2 := resp.CSS(`body`).CSS(`li`).CSS(`a[href=\/time\/]`).First().Text()
	if result2 != "2019-06-21" {
		t.Fatal("Response.CSS() failed.")
	}

	result3 := resp.CSS(`a`).First().Attr("href")
	if result3 != "/convenient/" {
		t.Fatal("Response.CSS().First().Attr() failed.")
	}

	result5 := resp.CSS(`a`).At(2).Text()
	if result5 != "南北" {
		t.Fatal("Response.CSS().At() failed.")
	}
}

func TestResponseEncoding(t *testing.T) {
	ts := newTestResponseServer()
	defer ts.Close()

	resp3, err := Get(ts.URL + "/latin1")
	if err != nil {
		t.Fatal(err)
	}
	resp3.Encoding("latin1")
	result3 := resp3.ReSubmatch(`<a href="/author/">(.*?)</a>`)
	if result3[0][0] != "..." {
		t.Fatal("Response latin1 failed.")
	}

	resp, err := Get(ts.URL + "/GBK")
	if err != nil {
		t.Fatal(err)
	}
	resp.Encoding("GBK")
	result1 := resp.ReSubmatch(`<a href="/author/">(.*?)</a>`)
	if result1[0][0] != "南北" {
		t.Fatal("Response GBK failed.")
	}

	resp2, err := Get(ts.URL + "/GB18030")
	if err != nil {
		t.Fatal(err)
	}
	resp2.Encoding("GB18030")
	result2 := resp2.ReSubmatch(`<a href="/author/">(.*?)</a>`)
	if result2[0][0] != "南北" {
		t.Fatal("Response GB18030 failed.")
	}
}
