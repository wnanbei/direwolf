package direwolf

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

type Request struct{

}

type Session struct{
	client  *http.Client
}
func (self *Session) Request(method string, url string, kwargs ...interface{}) {
	self.client = &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := self.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func (self *Session) Get(url string, kwargs ...interface{}){
	self.Request("Get", url, kwargs...)
}