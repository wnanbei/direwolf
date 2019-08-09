package direwolf

import (
	"bytes"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// Response is the response from request.
type Response struct {
	URL        string
	StatusCode int
	Proto      string
	Encoding   string
	Headers    http.Header
	Request    *RequestSetting
	content    []byte
	dom        *goquery.Document
}

// Content read bytes from Response.body.
func (resp *Response) Content() []byte {
	return resp.content
}

// Text decode content to string. You can specified encoding type. Such as GBK, GB18030,
// latin1. Default is UTF-8.
//
// If Response.content doesn`t exists, it will call Response.Content() at first.
func (resp *Response) Text(encoding ...string) string {
	var text = ""
	var encodingType = strings.ToUpper(resp.Encoding)

	if len(encoding) > 0 {
		encodingType = strings.ToUpper(encoding[0])
	}

	if resp.content == nil {
		resp.Content()
	}

	switch encodingType {
	case "UTF-8", "UTF8":
		text = string(resp.content)
	case "GBK":
		decodeBytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(resp.content)
		if err != nil {
			return ""
		}
		text = string(decodeBytes)
	case "GB18030":
		decodeBytes, err := simplifiedchinese.GB18030.NewDecoder().Bytes(resp.content)
		if err != nil {
			return ""
		}
		text = string(decodeBytes)
	case "LATIN1":
		decodeBytes, err := charmap.ISO8859_1.NewDecoder().Bytes(resp.content)
		if err != nil {
			return ""
		}
		text = string(decodeBytes)
	}

	return text
}

// Re extract required data with regexp.
// It return a slice of string.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) Re(queryStr string) []string {
	text := resp.Text()
	return regexp.MustCompile(queryStr).FindAllString(text, -1)
}

// ReSubmatch extract required data with regexp.
// It return a slice of string from FindAllStringSubmatch.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) ReSubmatch(queryStr string) [][]string {
	text := resp.Text()
	data := regexp.MustCompile(queryStr).FindAllStringSubmatch(text, -1)
	var subMatchResult [][]string
	for _, match := range data {
		if len(match) > 1 { // In case that query has no submatch part
			subMatchResult = append(subMatchResult, match[1:len(match)])
		}
	}
	return subMatchResult
}

// CSS is a method to extract data with css selector, it returns a CSSNodeList.
func (resp *Response) CSS(queryStr string) *CSSNodeList {
	content := bytes.NewReader(resp.Content())
	dom, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		return nil
	}
	resp.dom = dom

	newNodeList := []CSSNode{}
	resp.dom.Find(queryStr).Each(func(i int, selection *goquery.Selection) {
		newNode := CSSNode{selection: selection}
		newNodeList = append(newNodeList, newNode)
	})
	return &CSSNodeList{container: newNodeList}
}

// CSSNode is a container that stores single selected results
type CSSNode struct {
	selection *goquery.Selection
}

// Text return the text of the CSSNode. Only include stright children node text
func (node *CSSNode) Text() string {
	if node.selection != nil {
		var text string
		node.selection.Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "#text" {
				t := s.Text()
				if t != "" {
					text = text + t
				}
			}
		})
		return text
	}
	return ""
}

// TextAll return the text of the CSSNode. Include all children node text
func (node *CSSNode) TextAll() string {
	if node.selection != nil {
		return node.selection.Text()
	}
	return ""
}

// Attr return the attribute value of the CSSNode.
// You can set default value, if value isn`t exists, return default value.
func (node *CSSNode) Attr(attrName string, defaultValue ...string) string {
	var d, attrValue string
	if node.selection != nil {
		if len(defaultValue) > 0 {
			d = defaultValue[0]
		}
		attrValue = node.selection.AttrOr(attrName, d)
	}
	return attrValue
}

// CSSNodeList is a container that stores selected results
type CSSNodeList struct {
	container []CSSNode
}

// Text return a list of text. Only include stright children node text
func (nodeList *CSSNodeList) Text() (textList []string) {
	for _, node := range nodeList.container {
		text := node.Text()
		if text != "" {
			textList = append(textList, text)
		}
	}
	return
}

// TextAll return a list of text. Include all children node text
func (nodeList *CSSNodeList) TextAll() (textList []string) {
	for _, node := range nodeList.container {
		text := node.TextAll()
		if text != "" {
			textList = append(textList, text)
		}
	}
	return
}

// Attr return a list of attribute value
func (nodeList *CSSNodeList) Attr(attrName string, defaultValue ...string) (valueList []string) {
	for _, node := range nodeList.container {
		value := node.Attr(attrName, defaultValue...)
		if value != "" {
			valueList = append(valueList, value)
		}
	}
	return
}

// CSS return a CSSNodeList, so you can chain CSS
func (nodeList *CSSNodeList) CSS(queryStr string) *CSSNodeList {
	newNodeList := []CSSNode{}
	for _, node := range nodeList.container {
		node.selection.Find(queryStr).Each(func(i int, selection *goquery.Selection) {
			newNode := CSSNode{selection: selection}
			newNodeList = append(newNodeList, newNode)
		})
	}
	return &CSSNodeList{container: newNodeList}
}

// First return the frist cssNode of CSSNodeList.
// Return a empty cssNode if there is no cssNode in CSSNodeList
func (nodeList *CSSNodeList) First() *CSSNode {
	if len(nodeList.container) > 0 {
		return &nodeList.container[0]
	}
	return &CSSNode{}
}

// At return the cssNode of specified index position.
// Return a empty cssNode if there is no cssNode in CSSNodeList
func (nodeList *CSSNodeList) At(index int) *CSSNode {
	if len(nodeList.container) > index {
		return &nodeList.container[index]
	}
	return &CSSNode{}
}
