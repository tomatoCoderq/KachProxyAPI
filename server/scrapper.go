package server

import (
	"github.com/gocolly/colly"
)

//init scrapper here

func InitScrapper() *colly.Collector {
	scr := colly.NewCollector()
	return scr
}
