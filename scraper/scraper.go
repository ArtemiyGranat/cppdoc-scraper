package scraper

import (
	"cppreference-scraper/config"
	"cppreference-scraper/model"

	"fmt"

	"github.com/gocolly/colly/v2"
)

type Scraper struct {
	c      *colly.Collector
	config *config.Config
	pages  []model.Page
}

func New(config *config.Config) *Scraper {
	c := colly.NewCollector(
		colly.AllowedDomains(config.AllowedDomains...),
		colly.CacheDir(config.CacheDir),
	)
	c.UserAgent = "Googlebot"

	return &Scraper{
		c,
		config,
		[]model.Page{},
	}
}

func (s *Scraper) Scrape() []model.Page {
	s.setupCallbacks()
	s.c.Visit(s.config.StartingUrl)

	return s.pages
}

func (s *Scraper) setupCallbacks() {
	s.c.OnHTML(ContentSelector, s.extractPageData)
	s.c.OnHTML(LinkSelector, s.visitInternalLinks)

	s.c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
}
