package direwolf

// import (
// 	"net/http"
// )

// func Get() {
// 	http.Get(url)
// }

type Headers map[string]string
type Cookies map[string]string
type DataForm map[string]string
type Params map[string]string
type Data string

func Get(url string, kwargs ...interface{})  {
	session := Session{}
	session.Get(url, kwargs...)
}

func Post(url string, kwargs ...interface{})  {
	session := Session{}
	session.Post(url, kwargs...)
}