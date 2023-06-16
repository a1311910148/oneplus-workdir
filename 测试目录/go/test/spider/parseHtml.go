package spider

import (
	"strings"

	"golang.org/x/net/html"
)

func ParseHTML(htmlStr string) *html.Node {
	r := strings.NewReader(htmlStr)
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	return doc
}
