package direwolf

// Headers is request headers, as parameter in Request method.
// You can init it like this:
// 	headers := Headers{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type Headers map[string][]string

// Extend is aim to extend another map[string][]string to Headers.
// func (h Headers) Extend(src map[string][]string) {
// 	for key, value := range src {
// 		h[key] = value
// 	}
// }

// Params is url params you want to join to url, as parameter in Request method.
// You can init it like this:
// 	params := Params{
//  	"key1": {"value1", "value2"},
//  	"key2": {"value3"},
//  }
type Params map[string][]string

// DataForm is the form you want to post, as parameter in Request method.
// You can init it like this:
// 	df := DataForm{
// 		"key1": {"value1", "value2"},
// 		"key2": {"value3"},
//  }
type DataForm map[string][]string

// Cookies is request cookies, as parameter in Request method.
type Cookies map[string]string

// Data is data you want to post, as parameter in Request method.
type Data string
