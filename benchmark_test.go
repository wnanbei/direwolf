package direwolf

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/valyala/fasthttp"
)

func BenchmarkDirewolfGet(b *testing.B) {
	ts := newTestResponseServer()
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := Get(ts.URL)
		resp.Text()
	}
}

func BenchmarkGolangGet(b *testing.B) {
	ts := newTestResponseServer()
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get(ts.URL)
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
}

func BenchmarkFasthttpGet(b *testing.B) {
	ts := newTestResponseServer()
	defer ts.Close()
	b.ResetTimer()
	client := fasthttp.Client{}
	for i := 0; i < b.N; i++ {
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		req.SetRequestURI(ts.URL)
		client.Do(req, resp)
		req.Reset()
		resp.Reset()
	}
}
