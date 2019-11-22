package direwolf

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func BenchmarkDefaultGet(b *testing.B) {
	ts := newTestResponseServer()
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := Get(ts.URL)
		resp.Text()
	}
}

func BenchmarkGoGet(b *testing.B) {
	ts := newTestResponseServer()
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get(ts.URL)
		ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
}
