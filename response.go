package direwolf

import (
	"io"
	"io/ioutil"
	"log"
	"regexp"
)

// Response is the response from request.
type Response struct {
	URL        string
	StatusCode int
	Proto      string
	body       io.ReadCloser
	content    []byte
}

// Content read bytes from Response.body.
func (resp *Response) Content() []byte {
	if resp.content == nil {
		content, err := ioutil.ReadAll(resp.body)
		if err != nil {
			log.Fatal(err.Error())
		}
		resp.body.Close()
		resp.content = content
	}
	return resp.content
}

// Text decode content to string.
// if Response.content doesn`t exists, call Response.Content at first.
func (resp *Response) Text() string {
	var text = ""
	if resp.content == nil {
		text = string(resp.Content())
	} else {
		text = string(resp.content)
	}
	return text
}

// func (resp *Response) Css(query string) []string {

// }

// func (resp *Response) CssFirst(query string) string {

// }

// Re extract required data with regexp.
// It return a slice of string from FindAllString.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) Re(query string) []string {
	text := resp.Text()
	data := regexp.MustCompile(query).FindAllString(text, -1)
	return data
}

// ReSubmatch extract required data with regexp.
// It return a slice of string from FindAllStringSubmatch.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) ReSubmatch(query string) []string {
	text := resp.Text()
	data := regexp.MustCompile(query).FindAllStringSubmatch(text, -1)
	var subMatchData []string
	for _, match := range data {
		subMatchData = append(subMatchData, match[1])
	}
	return subMatchData
}

// func (resp *Response) ReFirst(query string) string {

// }
