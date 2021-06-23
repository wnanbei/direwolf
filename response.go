package direwolf

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// Response is the response from request.
type Response struct {
	URL           string
	StatusCode    int
	Proto         string
	Headers       http.Header
	Cookies       Cookies
	Request       *Request
	Content       []byte
	ContentLength int64
	encoding      string
	text          string
	dom           *goquery.Document
}

// Encoding can change and return the encoding type of response. Like this:
//   encoding := resp.Encoding("GBK")
// You can specified encoding type. Such as GBK, GB18030, latin1. Default is UTF-8.
// It will decode the content to string if you specified encoding type.
// It will just return the encoding type of response if you do not pass parameter.
func (resp *Response) Encoding(encoding ...string) string {
	if len(encoding) > 0 {
		resp.encoding = strings.ToUpper(encoding[0])
		resp.text = decodeContent(resp.encoding, resp.Content)
	}
	return resp.encoding
}

// Text return the text of Response. It will decode the content to string the first time
// it is called.
func (resp *Response) Text() string {
	if resp.text == "" {
		resp.text = decodeContent(resp.encoding, resp.Content)
	}
	return resp.text
}

// Re extract required data with regexp.
// It return a slice of string.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) Re(queryStr string) []string {
	text := resp.Text()
	return regexp.MustCompile(queryStr).FindAllString(text, -1)
}

// ReSubMatch extract required data with regexp.
// It return a slice of string from FindAllStringSubmatch.
// Every time you call this method, it will transcode the Response.content to text once.
// So please try to extract required data at once.
func (resp *Response) ReSubMatch(queryStr string) [][]string {
	text := resp.Text()
	data := regexp.MustCompile(queryStr).FindAllStringSubmatch(text, -1)
	var subMatchResult [][]string
	for _, match := range data {
		if len(match) > 1 { // In case that query has no sub match part
			subMatchResult = append(subMatchResult, match[1:])
		}
	}
	return subMatchResult
}

// CSS is a method to extract data with css selector, it returns a CSSNodeList.
func (resp *Response) CSS(queryStr string) *CSSNodeList {
	if resp.dom == nil { // New the dom if resp.dom not exists.
		text := strings.NewReader(resp.Text())
		dom, err := goquery.NewDocumentFromReader(text)
		if err != nil {
			return nil
		}
		resp.dom = dom
	}

	newNodeList := make([]CSSNode, 0)
	resp.dom.Find(queryStr).Each(func(i int, selection *goquery.Selection) {
		newNode := CSSNode{selection: selection}
		newNodeList = append(newNodeList, newNode)
	})
	return &CSSNodeList{container: newNodeList}
}

// Json can unmarshal json type response body to a struct.
func (resp *Response) Json(output interface{}) error {
	if err := jsoniter.Unmarshal(resp.Content, output); err != nil {
		return err
	}
	return nil
}

// JsonGet can get a value from json type response body with path.
func (resp *Response) JsonGet(path string) gjson.Result {
	return gjson.GetBytes(resp.Content, path)
}

// decodeContent decode the content with the encodingType. It just support
// UTF-8, GBK, GB18030, Lantin1 now.
func decodeContent(encodingType string, content []byte) (decodedText string) {
	var text = ""
	switch encodingType {
	case "UTF-8", "UTF8":
		text = string(content)
	case "GBK":
		decodeBytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(content)
		if err != nil {
			return ""
		}
		text = string(decodeBytes)
	case "GB18030":
		decodeBytes, err := simplifiedchinese.GB18030.NewDecoder().Bytes(content)
		if err != nil {
			return ""
		}
		text = string(decodeBytes)
	case "LATIN1":
		decodeBytes, err := charmap.ISO8859_1.NewDecoder().Bytes(content)
		if err != nil {
			return ""
		}
		text = string(decodeBytes)
	}
	return text
}

// CSSNode is a container that stores single selected results
type CSSNode struct {
	selection *goquery.Selection
}

// Text return the text of the CSSNode. Only include straight children node text
func (node *CSSNode) Text() string {
	if node.selection != nil {
		var text string
		node.selection.Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "#text" {
				t := s.Text()
				text = text + t
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

// Text return a list of text. Only include straight children node text
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
	newNodeList := make([]CSSNode, 0)
	for _, node := range nodeList.container {
		node.selection.Find(queryStr).Each(func(i int, selection *goquery.Selection) {
			newNode := CSSNode{selection: selection}
			newNodeList = append(newNodeList, newNode)
		})
	}
	return &CSSNodeList{container: newNodeList}
}

// First return the first cssNode of CSSNodeList.
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
