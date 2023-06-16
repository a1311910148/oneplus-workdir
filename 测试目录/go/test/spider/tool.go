package spider

// 爬虫
import (
	"fmt"
	"io/ioutil"
	"main/config"
	"net/http"

	"golang.org/x/net/html"
)

// 爬取网页内容
func SpiderPage() {
	// 获取url
	url := config.ToolUrl
	// 获取网页内容http
	resp := HttpGet(url)
	// 解析html
	doc := ParseHTML(resp)

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			// if n.Data == "title" {
			// 	fmt.Println("Title:", n.FirstChild.Data)
			// } else if n.Data == "p" {
			// 	fmt.Println("Paragraph:", n.FirstChild.Data)
			// }
			// if input
			if n.Data == "h2" {
				fmt.Println("input:", n.FirstChild.Attr)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
func HttpGet(url string) (content string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("http get error: " + err.Error())
		return ""
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("read body error: " + err.Error())
		return ""
	}
	return string(body)
}
