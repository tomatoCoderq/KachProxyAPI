package services

import (
	"github.com/tomatoCoderq/KachProxyAPI/models"
	"fmt"
	_ "fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tomatoCoderq/KachProxyAPI/internal"

	"github.com/gocolly/colly"
)

func scrappEverything() ([]*models.Play, error) {
	c := colly.NewCollector()

	playsScrapped := make([]*models.Play, 0)

	c.OnRequest(
		func(r *colly.Request) {
			log.Println("Visiting the page...", r.URL.String())
		})

	c.OnHTML("div.aff_el", func(e *colly.HTMLElement) {
		month := internal.RemoveDay(e.ChildText("div.day")) //obtaining month of specific play
		currentMonth := int(time.Now().Month())

		day, err := strconv.Atoi(e.ChildText("div.date")) //obtaining day of specific play
		if err != nil {
			log.Panicln("Error occured during casting from ascii to int")
		}

		link := e.ChildAttr("a", "href")
		var id string

		if strings.HasPrefix(link, "/affiche/detail/") {
			paresedURL, err := url.Parse(link)
			if err != nil {
				return
			}
			id = paresedURL.Query().Get("id")
		}

		sceneAndAuthor := e.ChildText("div.author")
		var author string
		var scene string

		//splitting scene and author
		if strings.HasPrefix(sceneAndAuthor, "Основная сцена") {
			scene = "Основная сцена"
			author = strings.TrimPrefix(sceneAndAuthor, "Основная сцена")
		} else if strings.HasPrefix(sceneAndAuthor, "Малая сцена") {
			scene = "Малая сцена"
			author = strings.TrimPrefix(sceneAndAuthor, "Малая сцена")
		}

		if month == 1 && day >= time.Now().Day() {
			playsScrapped = append(playsScrapped, &models.Play{
				Id:        id,
				Scene:     scene,
				Author:    author,
				Name:      e.ChildText("div.name"),
				Genre:     e.ChildText("div.genre"),
				AgeRating: strings.TrimSuffix(e.ChildText("div.age_rating"), "+"),
				Month:     month,
				Day:       day,
			})
		} else if month >= currentMonth && day >= time.Now().Day() {
			playsScrapped = append(playsScrapped, &models.Play{
				Id:        id,
				Scene:     scene,
				Author:    author,
				Name:      e.ChildText("div.name"),
				Genre:     e.ChildText("div.genre"),
				AgeRating: strings.TrimSuffix(e.ChildText("div.age_rating"), "+"),
				Month:     month,
				Day:       day,
			})
		}
	})

	c.OnScraped(func(e *colly.Response) {
		log.Println("Scrapping finished")
	})

	err := c.Visit("https://teatrkachalov.ru/affiche/base/")
	if err != nil {
		return nil, err
	}

	return playsScrapped, nil
}

type PlaysService struct {
	scrapper *colly.Collector
}

func NewPlaysService(scrapper *colly.Collector) *PlaysService {
	return &PlaysService{
		scrapper: scrapper,
	}
}

func (ps *PlaysService) GetAllPlays(scene string, author string, name string, age_rating string, month string) ([]*models.Play, error) {
	playsScrapped, err := scrappEverything()
	if err != nil {
		return nil, err
	}

	
	filteredPlays := make([]*models.Play, 0)
	
	for _, play := range playsScrapped {
		if scene != "" && play.Scene != scene {
			continue
		}
		if author != "" && !strings.HasSuffix(play.Author, author) {
			continue
		}
		if name != "" && play.Name != name {
			continue
		}
		if age_rating != "" && play.AgeRating != age_rating {
			continue
		}
		if month != "" && fmt.Sprintf("%d", play.Month) != month {
			continue
		}
		
		filteredPlays = append(filteredPlays, play)
	}
	
	fmt.Println(len(filteredPlays))
	return filteredPlays, nil
}

func (ps *PlaysService) GetPlayById(id string) (*models.Play, error) {
	playsScrapped, err := scrappEverything()
	if err != nil {
		return nil, err
	}

	for _, play := range playsScrapped {
		if play.Id == id {
			return play, nil
		}
	}

	return nil, nil
}
