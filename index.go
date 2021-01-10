package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type a struct {
	Aa  string      `json:"aa,omitempty"`
	Ab  string      `json:"ab,omitempty"`
	BB  int         `json:"bb,omitempty"`
	CcD interface{} `json:"cc_d,omitempty"`
}

func main() {
	c := colly.NewCollector(
	// colly.Debugger(&debug.LogDebugger{}),
	)
	c.OnHTML(".info_list li .txt_area a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnHTML(".article-tit", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	// c.OnScraped()
	c.Visit("https://m.cnbeta.com/")

}
