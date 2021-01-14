package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	url := "http://nghiahsgs.com"
	c := colly.NewCollector()

	var listLink []string
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// fmt.Println("link", link)
		listLink = append(listLink, link)
	})
	c.OnRequest(func(e *colly.Request) {
		fmt.Println(e.URL)
	})
	c.Visit(url)

	fmt.Println("listLink", listLink)
}
