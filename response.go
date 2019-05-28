package direwolf

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// Response is the response from request.
type Response struct {
	URL        string
	StatusCode int
	Proto      string
	body       io.ReadCloser
	content    []byte
	dom        *goquery.Document
	Request    *Request
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

// CSS is a api to goquery, it returns a goquery.Selection object.
// so you can totally use the api from goquery, like Find().
func (resp *Response) CSS(query string) *goquery.Selection {
	content := bytes.NewReader(resp.Content())
	dom, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.dom = dom
	queryResult := resp.dom.Find(query)
	return queryResult
}

// CSSFirst return the first node text from query result.
func (resp *Response) CSSFirst(query string) string {
	queryResult := resp.CSS(query)
	return queryResult.First().Text()
}

// Re extract required data with regexp.
// It return a slice of string from FindAllString.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) Re(query string) []string {
	text := resp.Text()
	queryResult := regexp.MustCompile(query).FindAllString(text, -1)
	return queryResult
}

// ReSubmatch extract required data with regexp.
// It return a slice of string from FindAllStringSubmatch.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) ReSubmatch(query string) []string {
	text := resp.Text()
	data := regexp.MustCompile(query).FindAllStringSubmatch(text, -1)
	var subMatchResult []string
	for _, match := range data {
		if len(match) > 1 { // In case that query has no submatch part
			subMatchResult = append(subMatchResult, match[1])
		}
	}
	return subMatchResult
}

// func (resp *Response) ReFirst(query string) string {

// }
