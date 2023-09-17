package scraper

import (
	"cppreference-scraper/model"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func (s *Scraper) extractPageData(e *colly.HTMLElement) {
	page := model.Page{
		Title:       e.ChildText(TitleSelector),
		Description: extractDescription(e),
		Headers:     joinLines(e, HeadersSelector),
		Signature:   joinLines(e, SignatureSelector),
		Example:     e.ChildText(ExampleSelector),
	}
	s.pages = append(s.pages, page)
}

func (s *Scraper) visitInternalLinks(e *colly.HTMLElement) {
	href := e.Attr("href")
	hasVisited, err := s.c.HasVisited(href)
	if err != nil {
		fmt.Println(err)
	}
	if !hasVisited && strings.HasPrefix(href, "/w/cpp") {
		e.Request.Visit(e.Attr("href"))
	}
}

func extractDescription(e *colly.HTMLElement) string {
	description := ""
	e.ForEach(DescriptionDelimSelector, func(_ int, el *colly.HTMLElement) {
		precedingParagraphs := el.DOM.PrevAllFiltered(DescriptionSelector)
		for i := precedingParagraphs.Length() - 1; i >= 0; i-- {
			paragraph := precedingParagraphs.Get(i)
			description += goquery.NewDocumentFromNode(paragraph).Text()
			if i != 0 {
				description += "\n"
			}
		}
	})
	return description
}