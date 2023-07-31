package main

import (
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a.js_link h2", func(e *colly.HTMLElement) {
		title := e.Text
		link, exists := e.DOM.Parent().Attr("href")
		if exists {
			log.Println(title + " (" + link + ")")
		}
	})

	c.Visit("https://qz.com/latest")
}
