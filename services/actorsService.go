package services

import (
	// "KachProxyAPI/pkg/models"
	"KachProxyAPI/models"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type ActorsService struct {
	scrapper *colly.Collector
}

func NewActorsService(scrapper *colly.Collector) *ActorsService {
	return &ActorsService{
		scrapper: scrapper,
	}
}

//TODO: Page dynamically updates so I need to use some tool to fetch the page
func (as *ActorsService) GetAllActors() ([]*models.Actor, error) {
	c := colly.NewCollector()

	actors := make([]*models.Actor, 0)

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

		actors = append(actors, &models.Actor{
			Id:   id,
			Name: e.ChildText("div.name"),
		})
	})

	c.OnScraped(func(e *colly.Response) {
		log.Println("Scrapping finished")
	})

	return actors, c.Visit("https://teatrkachalov.ru/troupe/")
}

func (as *ActorsService) GetActorById(id string) (*models.Actor, error) {
	if id == "" {
		return nil, fmt.Errorf("actor id is empty")
	}

	actors, err := as.GetAllActors()
	if err != nil {
		return nil, err
	}
	for _, actor := range actors {
		if actor.Id == id {
			return actor, nil
		}
	}
	return nil, nil
}