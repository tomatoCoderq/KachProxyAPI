package services

import (
	"KachProxyAPI/models"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type MusiciansService struct {
	scrapper *colly.Collector
}

func NewMusiciansService(scrapper *colly.Collector) *MusiciansService {
	return &MusiciansService {
		scrapper: scrapper,
	}
}


func (as *MusiciansService) GetAllMusicians() ([]*models.Musicians, error) {
	c := colly.NewCollector()

	musicians := make([]*models.Musicians, 0)

	c.OnRequest(
		func(r *colly.Request) {
			log.Println("Visiting the page...", r.URL.String())
		})

	c.OnHTML("div.actor", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		var id string

		if strings.HasPrefix(link, "detail/") {
			paresedURL, err := url.Parse(link)
			if err != nil {
				return
			}
			id = paresedURL.Query().Get("id")
		}

		musicians = append(musicians, &models.Musicians{
			Id:   id,
			Name: e.ChildText("div.name"),
		})
	})

	c.OnScraped(func(e *colly.Response) {
		log.Println("Scrapping finished")
	})

	return musicians, c.Visit("https://teatrkachalov.ru/musicians/")
}

func (as *MusiciansService) GetMusicianById(id string) (*models.Musicians, error) {
	if id == "" {
		return nil, fmt.Errorf("musician id is empty")
	}

	musicians, err := as.GetAllMusicians()
	if err != nil {
		return nil, err
	}
	for _, musician := range musicians {
		if musician.Id == id {
			return musician, nil
		}
	}
	return nil, nil
}