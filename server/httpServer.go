package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/spf13/viper"

	"KachProxyAPI/controllers"
	// "KachProxyAPI/pkg/models"
	"KachProxyAPI/services"
)


type HttpServer struct {
	router *gin.Engine
	config *viper.Viper

}

func InitHttpServer(config *viper.Viper, scrapper *colly.Collector) HttpServer{
	playsService := services.NewPlaysService(scrapper)
	actorsService := services.NewActorsService(scrapper)
	musiciansService := services.NewMusiciansService(scrapper)

	playsController := controllers.NewPlaysController(playsService)
	actorsController := controllers.NewActorsController(actorsService)
	musiciansController := controllers.NewMusiciansController(musiciansService)

	router := gin.Default()
	router.GET("/plays", playsController.GetAllPlays)
	router.GET("/plays/:id", playsController.GetPlayById)
	router.GET("/actors", actorsController.GetAllActors)
	router.GET("/actors/:id", actorsController.GetActorById)
	router.GET("/musicians", musiciansController.GetAllMusicians)
	router.GET("/musicians/:id", musiciansController.GetMusicianById)

	return HttpServer{
		router: router,
		config:  config,
	}

}

func (hs HttpServer)Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatal("Error during running server")
	}
}

