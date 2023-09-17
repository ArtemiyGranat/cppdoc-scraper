package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	LinkSelector             = "#content a[href]"
	ContentSelector          = "div[id=content]"
	TitleSelector            = "h1[id=firstHeading]"
	DescriptionDelimSelector = "table.toc"
	DescriptionSelector      = "p, ul, li"
	HeadersSelector          = ".t-dcl-begin .t-dsc-header"
	SignatureSelector        = ".t-dcl"
	ExampleSelector          = "div.cpp.source-cpp"
)

func joinLines(e *colly.HTMLElement, selector string) string {
	return strings.ReplaceAll(strings.TrimSpace(e.ChildText(selector)), "\n", "")
}